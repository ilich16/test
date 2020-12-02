<template>
  <v-app id="inspire">
    <v-container class="fill-height" fluid>
      <v-row align="center" justify="center">
        <v-col cols="12" sm="8" md="4">
          <v-card class="elevation-12">
            <v-toolbar color="deep-purple accent-4" dark flat>
              <v-toolbar-title>Iniciar sesión</v-toolbar-title>
            </v-toolbar>
            <v-card-text>
              <v-form ref="form" v-model="valid" lazy-validation>
                <v-text-field
                  color="deep-purple accent-4"
                  v-model="username"
                  label="Nombre de usuario"
                  :rules="inputRules"
                  prepend-icon="mdi-account"
                  type="text"
                  required
                ></v-text-field>

                <v-text-field
                  color="deep-purple accent-4"
                  v-model="password"
                  :append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
                  :rules="inputRules"
                  :type="showPassword ? 'text' : 'password'"
                  label="Contraseña"
                  counter
                  prepend-icon="mdi-lock"
                  @click:append="showPassword = !showPassword"
                  required
                ></v-text-field>
              </v-form>
            </v-card-text>
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn
                text
                color="deep-purple accent-4"
                :loading="loadingButton"
                :disabled="loadingButton"
                @click="Login"
                >Ingresar</v-btn
              >
            </v-card-actions>
          </v-card>
        </v-col>
      </v-row>
      <result-dialog
        v-bind:responseCode="responseCode"
        v-bind:resultDialog="resultDialog"
        v-on:close-dialog="resultDialog = $event"
      ></result-dialog>
    </v-container>
  </v-app>
</template>

<script>
import ResultDialog from "./ResultDialog";

const axios = require("axios").default;
const apiURL = "https://testing06.com:8080";

export default {
  name: "login-form",

  components: {
    ResultDialog,
  },

  data: function() {
    return {
      resultDialog: false,
      loader: null,
      loadingButton: false,
      valid: true,
      username: "",
      password: "",
      dialogMessage: "",
      responseCode: 0,
      showPassword: false,
      inputRules: [(v) => !!v || "Campo requerido"],
    };
  },

  methods: {
    Login: function() {
      if (this.$refs.form.validate()) {
        this.loadingButton = true;
        axios
          .post(apiURL + "/api/v1/sw1/login", {
            username: this.username,
            password: this.password,
          })
          .then((response) => {
            this.responseCode = response.status;
            this.resultDialog = true;
            this.loadingButton = false;
          })
          .catch((error) => {
            this.responseCode = error.response.status;
            this.resultDialog = true;
            this.loadingButton = false;
          });
      }
    },
  },
};
</script>
