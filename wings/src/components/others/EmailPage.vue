<template>
    <div class="ui segment">
        <div v-if="!dedaloRequested">
            <div class="ui big left icon input">
                <input v-model="authEmail" type="email" placeholder="Insert your email">
                <i class="talk icon"></i>
            </div>
            <button v-on:click="getCode()" class="ui big button auth-code-cont">Get Code</button>
            <div v-if="errors.badMail" class="ui tiny icon negative message">
                <i class="remove icon"></i>
                <div class="content">
                    <div class="header">
                        Error sending verification code
                    </div>
                    <p>Something gone wrong sending code to your email</p>
                </div>
            </div>
            <div v-if="codeRequested" class="auth-code-cont">
                <div class="ui big left icon input">
                    <input v-model="authCode" type="number" placeholder="Insert your code">
                    <i class="braille icon"></i>
                </div>
            </div>
            <div class="ui divider"></div>
            <button v-on:click="execLogin()" :disabled="isDisabled()" class="big ui green button">
                Start Navigate
            </button>
        </div>
        <div v-if="dedaloRequested">
            <div v-if="!authorized" class="ui active centered inline text loader">Authorization in progress...</div>
            <div v-if="authorized" class="ui icon positive message">
                <i class="check icon"></i>
                <div class="content">
                    <div class="header">
                        You are successfully authenticated
                    </div>
                    <p>In a few seconds you will be redirected...</p>
                </div>
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
            var authorized = false
            var codeRequested = false
            var dedaloRequested = false
            var badMail = false
            var badCode = false

            return {
                authorized: authorized,
                codeRequested: codeRequested,
                dedaloRequested: dedaloRequested,
                authEmail: '',
                authCode: '',
                errors: {
                    badMail: badMail,
                    badCode: badCode
                }
            }
        },
        methods: {
            isDisabled() {
                return this.authEmail.length == 0 || this.authCode.length == 0
            },
            getCode() {
                this.errors.badMail = false
                if (!this.authEmail.indexOf('@') == -1) {
                    this.errors.badMail = true
                }
                var params = this.extractParams()

                // make request to wax
                var url = this.createWaxURL(this.authEmail, params, 'email')

                // get user id
                this.$http.get(url).then(responseAuth => {
                    this.codeRequested = true
                    this.authCode = responseAuth.body.password || ''
                }, response => {
                    this.codeRequested = false
                    this.errors.badMail = true
                    console.error(response)
                });
            },
            execLogin() {
                this.dedaloRequested = true
                this.authorized = false
                this.errors.badCode = false

                // exec dedalo login
                this.doDedaloLogin({
                    id: this.authEmail,
                    password: this.authCode || ''
                }, responseDedalo => {
                    if (responseDedalo.body.clientState == 1) {
                        this.authorized = true
                        setTimeout(function () {
                            // open redir url
                            window.location.replace(this.$root.$options.hotspot.preferences
                                .captive_redir)
                        }.bind(this), 2500)
                    } else {
                        this.authorized = false
                        this.errors.badCode = true
                    }
                }, error => {
                    this.authorized = false
                    console.error(error)
                })
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
</style>