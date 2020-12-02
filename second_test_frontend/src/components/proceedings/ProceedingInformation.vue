<template>
  <div>
    <div class="pa-4">
      <v-card class="mx-auto" max-width="600">
        <v-card-text>
          <p class="display-1 text--primary">
            Buscar acta
          </p>
          <div class="text--primary">
            ¿Estás buscando una acta en específico? Por favor ingresa el código
            de la acta y luego presiona en el botón para cargar sus datos
            correspondientes.
          </div>
        </v-card-text>
        <v-form ref="form" v-model="valid" :lazy-validation="lazy"></v-form>
        <v-autocomplete
          v-model="proceedingCode"
          class="mx-4"
          :items="totalProceedings"
          color="deep-purple accent-4"
          item-text="code"
          label="Acta"
        ></v-autocomplete>
        <v-card-actions>
          <v-btn
            text
            color="deep-purple accent-4"
            @click="getProceedingResults"
          >
            Cargar datos
          </v-btn>
        </v-card-actions>
      </v-card>
      <div class="pa-4">
        <v-card class="mx-auto" v-if="showProceedingValues" max-width="1000">
          <v-container>
            <v-row>
              <v-col cols="12" sm="6">
                <b>Código de acta:</b> {{ proceedingValues.codigo }}
              </v-col>
              <v-col cols="12" sm="6">
                <b>Recinto:</b> {{ proceedingValues.recinto }}
              </v-col>
              <v-col cols="12" sm="6">
                <b>Mesa:</b> {{ proceedingValues.mesa }}
              </v-col>
              <v-col cols="12" sm="6">
                <b>Cir. Uninominal:</b>
                {{ proceedingValues.circunscripcion }}
              </v-col>
              <v-col cols="12" sm="6">
                <b>Departamento:</b> {{ proceedingValues.departamento }}
              </v-col>
              <v-col cols="12" sm="6">
                <b>Provincia:</b> {{ proceedingValues.provincia }}
              </v-col>
              <v-col cols="12" sm="6">
                <b>Municipio:</b> {{ proceedingValues.municipio }}
              </v-col>
              <v-col cols="12" sm="6">
                <b>Localidad:</b> {{ proceedingValues.localidad }}
              </v-col>
            </v-row>
          </v-container>
          <v-simple-table>
            <template v-slot:default>
              <thead>
                <tr>
                  <th class="text-left">Partido político</th>
                  <th class="text-left">Presidente/a</th>
                  <th class="text-left">Diputado/a Cir. Uninominal</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="item in proceedingTable" :key="item.ciudad">
                  <td>{{ item.partido }}</td>
                  <td>{{ item.votosPresidente }}</td>
                  <td>{{ item.votosDiputado }}</td>
                </tr>
              </tbody>
            </template>
          </v-simple-table>
          <v-card-actions>
            <v-btn
              outlined
              color="deep-purple accent-4"
              @click="generateReport"
            >
              Generar reporte
            </v-btn>
            <v-btn
              outlined
              color="deep-purple accent-4"
              @click="downloadProceeding"
            >
              Descargar acta
            </v-btn>
          </v-card-actions>
          <h4 class="pa-4">
            Puedes obtener la ubicación de donde fue cargada la acta:
          </h4>
          <v-btn
            class="mx-4"
            outlined
            color="deep-purple accent-4"
            @click="generateLocation"
          >
            Obtener ubicación
          </v-btn>
          <h4 class="pa-4">
            Esta acta fue cargada desde el siguiente dispositivo:
            {{ proceedingValues.dispositivo }}
          </h4>
        </v-card>
      </div>
    </div>
  </div>
</template>

<script>
import fileDownload from "js-file-download";

const axios = require("axios").default;
const apiURL = "https://testing06.com:8080";

export default {
  data: function() {
    return {
      valid: false,
      lazy: false,
      proceedingCode: "",
      proceedingCodeRules: [
        (v) => v.length == 5 || "El código de acta debe ser de 5 dígitos.",
        (v) => !!v || "Este campo es requerido.",
      ],
      proceedingValues: {
        codigo: 0,
        recinto: "",
        mesa: 0,
        circunscripcion: 0,
        ciudad: "",
        provincia: "",
        municipio: "",
        localidad: "",
        partido: "",
        votosPresidente: 0,
        votosDiputado: 0,
        dispositivo: "",
        ubicacion: "",
      },
      proceedingTable: [],
      totalProceedings: [],
      showProceedingValues: false,
    };
  },

  mounted() {
    axios
      .get(apiURL + "/api/v1/sw1/available-proceedings")
      .then((response) => {
        this.totalProceedings = response.data;
      })
      .catch((error) => {
        console.log(error);
      });
  },

  methods: {
    getProceedingResults: function() {
      axios
        .get(apiURL + "/api/v1/sw1/proceeding-votes/" + this.proceedingCode)
        .then((response) => {
          this.proceedingValues = response.data[0];
          this.proceedingTable = response.data;
          this.showProceedingValues = true;
        });
    },
    generateLocation: function() {
      window.open(
        "https://www.google.com/maps/search/?api=1&query=" +
          this.proceedingValues.ubicacion
      );
    },
    generateReport: function() {
      axios
        .post(
          apiURL + "/api/v1/sw1/generate-report-proceeding",
          this.proceedingTable,
          {
            responseType: "blob",
          }
        )
        .then((response) => {
          var today = new Date();
          var date =
            today.getFullYear() +
            "-" +
            (today.getMonth() + 1) +
            "-" +
            today.getDate();
          var time =
            today.getHours() +
            ":" +
            today.getMinutes() +
            ":" +
            today.getSeconds();
          var dateTime = date + " " + time;
          fileDownload(response.data, dateTime + ".xlsx");
          console.log(response);
        })
        .catch((error) => {
          console.log(error);
        });
    },
    downloadProceeding: function() {
      axios
        .post(
          apiURL +
            "/api/v1/sw1/download-proceeding/" +
            this.proceedingValues.codigo,
          this.proceedingTable,
          {
            responseType: "blob",
          }
        )
        .then((response) => {
          fileDownload(response.data, this.proceedingValues.codigo + ".jpg");
        })
        .catch((error) => {
          console.log(error);
        });
    },
  },
};
</script>
