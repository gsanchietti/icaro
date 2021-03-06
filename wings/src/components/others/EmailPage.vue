<template>
    <div class="ui segment form">
        <div v-if="!dedaloRequested">
            <div v-if="!codeRequested" class="inline field" v-bind:class="{ error: errors.badInput }">
                <label>Email</label>
                <div class="ui big left icon input">
                    <input v-model="authEmail" type="email" :placeholder="$t('email.insert_email')">
                    <i class="mail icon"></i>
                </div>
            </div>
            <button v-if="!codeRequested" v-on:click="getCode()" class="ui big button request-code">{{ $t("email.have_code") }}</button>
            <button v-if="!codeRequested" v-on:click="getCode(true)" class="ui big button">{{ $t("email.get_code") }}</button>
            <div v-if="errors.badMail" class="ui tiny icon negative message">
                <i class="remove icon"></i>
                <div class="content">
                    <div class="header">
                        {{ $t("email.error_code") }}
                    </div>
                    <p>{{ $t("email.error_code_sub") }}</p>
                </div>
            </div>
            <div v-if="codeRequested">
                <div class="inline field">
                    <label>{{ $t("email.code") }}</label>
                    <div class="ui big left icon input">
                        <input v-model="authCode" type="number" :placeholder="$t('email.insert_your_code')">
                        <i class="braille icon"></i>
                    </div>
                </div>
            </div>
            <div class="ui divider"></div>
            <button v-on:click="execLogin()" :disabled="isDisabled()" class="big ui green button">
                {{ $t("email.start_navigate") }}
            </button>
            <div v-if="authReset && resetDone != 'true'" class="ui divider"></div>
            <button v-on:click="getCode(true)" v-if="authReset && resetDone != 'true'" class="ui red button">
                {{ $t("email.reset_code") }}
            </button>
        </div>
        <div v-if="dedaloRequested">
            <div v-if="!authorized && !errors.dedaloError" class="ui active centered inline text loader">{{ $t("email.auth_progress") }}...</div>
            <div v-if="authorized" class="ui icon positive message">
                <i class="check icon"></i>
                <div class="content">
                    <div class="header">
                        {{ $t("email.auth_success") }}
                    </div>
                    <p>{{ $t("email.auth_success_sub") }}...</p>
                </div>
            </div>
            <div v-if="errors.dedaloError" class="ui icon negative message">
                <i class="remove icon"></i>
                <div class="content">
                    <div class="header">
                        {{ $t("email.auth_error") }}
                    </div>
                    <p>{{ $t("email.auth_error_sub") }}</p>
                </div>
            </div>
            <div v-if="authorized">
                <h3>{{ $t("login.disclaimer_marketing") }}</h3>
                <div class="inline field">
                    <textarea readonly class="text-center" v-model="hotspot.disclaimers.marketing_use"></textarea>
                </div>
                <button v-on:click="deleteInfo()" class="ui big button red">{{ $t("login.decline") }}</button>
                <button v-on:click="accept()" class="ui big button green">{{ $t("login.accept") }}</button>
            </div>
        </div>
    </div>
</template>

<script>
    import AuthMixin from './../../mixins/auth';
    import {
        setTimeout
    } from 'timers';
    export default {
        name: 'EmailPage',
        mixins: [AuthMixin],
        data() {
            var params = this.extractParams()

            this.getPreferences(params, success => {
                this.$parent.hotspot.disclaimers = success.body.disclaimers
                this.$root.$options.hotspot.disclaimers = success.body.disclaimers
                this.hotspot.disclaimers = success.body.disclaimers
            }, error => {
                console.error(error)
            })

            return {
                authorized: false,
                codeRequested: this.$route.query.code || false,
                dedaloRequested: false,
                authEmail: this.$route.query.email || '',
                authCode: this.$route.query.code || '',
                authReset: this.$route.query.code || false,
                userId: this.$route.query.user || 0,
                resetDone: false,
                errors: {
                    badMail: false,
                    badCode: false,
                    dedaloError: false,
                    badInput: false
                },
                hotspot: {
                    disclaimers: this.$root.$options.hotspot.disclaimers
                },
            }
        },
        methods: {
            isDisabled() {
                return this.authEmail.length == 0 || this.authCode.length == 0
            },
            getCode(reset) {
                this.errors.badMail = false
                if (this.authEmail.indexOf('@') == -1) {
                    this.errors.badInput = true
                    return
                }
                var params = this.extractParams()

                // make request to wax
                var url = this.createWaxURL(this.authEmail, params, 'email', reset)

                // get user id
                this.$http.get(url).then(responseAuth => {
                    this.authReset = responseAuth.body.exists
                    this.resetDone = responseAuth.body.reset
                    this.userId = responseAuth.body.user_db_id

                    // check if user already exists
                    if (this.authReset && !(this.resetDone && this.resetDone == 'true')) {
                        this.codeRequested = true
                    } else {
                        // open temp session for the user
                        this.doTempSession(this.authEmail, responseTmp => {
                            this.codeRequested = true
                        }, error => {
                            this.codeRequested = false
                            this.errors.badMail = true
                            console.error(error)
                        })
                    }
                }, error => {
                    this.codeRequested = false
                    this.errors.badMail = true
                    console.error(error)
                });
            },
            execLogin() {
                this.dedaloRequested = true
                this.authorized = false
                this.errors.dedaloError = false
                this.errors.badCode = false

                // exec logout
                this.doDedaloLogout(responseDedaloLogout => {
                    // exec dedalo login
                    this.doDedaloLogin({
                        id: this.authEmail,
                        password: this.authCode || ''
                    }, responseDedalo => {
                        if (responseDedalo.body.clientState == 1) {
                            this.authorized = true
                            this.errors.dedaloError = false
                        } else {
                            this.authorized = false
                            this.errors.dedaloError = true
                            this.errors.badCode = true
                        }
                    }, error => {
                        this.authorized = false
                        this.errors.dedaloError = true
                        console.error(error)
                    })
                }, error => {
                    this.authorized = false
                    this.errors.dedaloError = true
                    console.error(error)
                })
            },
            deleteInfo() {
                // extract code and state
                var params = this.extractParams()
                this.deleteMarketingInfo(this.userId, params, function (success) {
                    this.accept()
                }, function (error) {
                    console.error(error)
                    if (error.status == 404) {
                        this.accept()
                    }
                })
            },
            accept() {
                // open redir url
                window.location.replace(this.$root.$options.hotspot.preferences
                    .captive_1_redir)
            }
        }
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
    h1,
    h2 {
        font-weight: normal;
    }

    ul {
        list-style-type: none;
        padding: 0;
    }

    li {
        display: inline-block;
        margin: 0 10px;
    }

    a {
        color: #42b983;
    }

    .auth-code-cont {
        margin-top: 15px !important;
        margin: 0;
    }

    .request-code {
        margin-bottom: 10px !important;
    }
</style>