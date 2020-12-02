<template>
  <div>
    <v-card flat>
      <v-card class="ma-4">
        <ve-bar
          class="px-4"
          :data="sucre.chartData"
          :extend="sucre.chartExtend"
          :legend-visible="false"
        ></ve-bar>
      </v-card>
      <v-card class="ma-4">
        <v-simple-table>
          <template v-slot:default>
            <thead>
              <tr>
                <th class="text-left">
                  Candidato a presidente y vicepresidente
                </th>
                <th class="text-left">Partido</th>
                <th class="text-left">Votos</th>
                <th class="text-left">%</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="item in sucre.results" :key="item.candidato">
                <td>{{ item.candidato }}</td>
                <td>{{ item.partido }}</td>
                <td>{{ item.votos }}</td>
                <td>{{ item.porcentaje }}</td>
              </tr>
            </tbody>
          </template>
        </v-simple-table>
        <v-card-actions>
          <v-btn
            outlined
            color="deep-purple accent-4"
            @click="getReportsucre"
          >
            Generar reporte
          </v-btn>
        </v-card-actions>
      </v-card>
      <v-expansion-panels class="px-4">
        <v-expansion-panel>
          <v-expansion-panel-header
            >Resultados por recinto</v-expansion-panel-header
          >
          <v-expansion-panel-content>
            <v-card class="mx-auto" max-width="600">
              <v-card-text>
                <p class="display-1 text--primary">
                  Buscar recinto
                </p>
                <div class="text--primary">
                  ¿Estás buscando un recinto en específico? Busca en la
                  siguiente lista de recintos disponibles y luego presiona en el
                  botón para cargar sus datos correspondientes.
                </div>
              </v-card-text>
              <v-autocomplete
                v-model="sucre.precinctValue"
                class="mx-4"
                :items="sucre.precincts"
                color="deep-purple accent-4"
                item-text="nombre"
                label="Recinto"
              ></v-autocomplete>
              <v-card-actions>
                <v-btn
                  text
                  color="deep-purple accent-4"
                  @click="getPrecinctResultsucre"
                >
                  Cargar datos
                </v-btn>
              </v-card-actions>
            </v-card>
            <v-card v-if="sucre.showPrecinctResult" class="ma-4">
              <v-simple-table>
                <template v-slot:default>
                  <thead>
                    <tr>
                      <th class="text-left">
                        Candidato a presidente y vicepresidente
                      </th>
                      <th class="text-left">Partido</th>
                      <th class="text-left">Votos</th>
                      <th class="text-left">%</th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr
                      v-for="item in sucre.precinctResult"
                      :key="item.candidato"
                    >
                      <td>{{ item.candidato }}</td>
                      <td>{{ item.partido }}</td>
                      <td>{{ item.votos }}</td>
                      <td>{{ item.porcentaje }}</td>
                    </tr>
                  </tbody>
                </template>
              </v-simple-table>
              <v-card-actions>
                <v-btn
                  outlined
                  color="deep-purple accent-4"
                  @click="getReportRecinctsucre"
                >
                  Generar reporte
                </v-btn>
              </v-card-actions>
            </v-card>
          </v-expansion-panel-content>
        </v-expansion-panel>
        <v-expansion-panel>
          <v-expansion-panel-header
            >Resultados por circunscripción</v-expansion-panel-header
          >
          <v-expansion-panel-content>
            <v-card class="mx-auto" max-width="600">
              <v-card-text>
                <p class="display-1 text--primary">
                  Buscar circunscripción
                </p>
                <div class="text--primary">
                  ¿Estás buscando una circunscripción en específico? Busca en la
                  siguiente lista de circunscripciones disponibles y luego
                  presiona en el botón para cargar sus datos correspondientes.
                </div>
              </v-card-text>
              <v-autocomplete
                v-model="sucre.districtValue"
                class="mx-4"
                :items="sucre.districts"
                color="deep-purple accent-4"
                label="Circunscripción"
              ></v-autocomplete>
              <v-card-actions>
                <v-btn
                  text
                  color="deep-purple accent-4"
                  @click="getDistrictResultsucre"
                >
                  Cargar datos
                </v-btn>
              </v-card-actions>
            </v-card>
            <v-card v-if="sucre.showDistrictResult" class="ma-4">
              <v-simple-table>
                <template v-slot:default>
                  <thead>
                    <tr>
                      <th class="text-left">
                        Candidato a presidente y vicepresidente
                      </th>
                      <th class="text-left">Partido</th>
                      <th class="text-left">Votos</th>
                      <th class="text-left">%</th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr
                      v-for="item in sucre.districtResult"
                      :key="item.candidato"
                    >
                      <td>{{ item.candidato }}</td>
                      <td>{{ item.partido }}</td>
                      <td>{{ item.votos }}</td>
                      <td>{{ item.porcentaje }}</td>
                    </tr>
                  </tbody>
                </template>
              </v-simple-table>
              <v-card-actions>
                <v-btn
                  outlined
                  color="deep-purple accent-4"
                  @click="getReportsucreDistrict"
                >
                  Generar reporte
                </v-btn>
              </v-card-actions>
            </v-card>
          </v-expansion-panel-content>
        </v-expansion-panel>
      </v-expansion-panels>
    </v-card>
  </div>
</template>

<script>
import VeBar from "v-charts/lib/bar.common";
import fileDownload from "js-file-download";

const COLOR_LIST = [
  "#e3002d",
  "#f979ac",
  "#018b7c",
  "#e3002d",
  "#00c2f7",
  "#e4bc7f",
  "#00b8ac",
  "#e95c15",
  "#01319f",
];

const partidos = [
  "MAS-IPSP",
  "C.C.",
  "FPV",
  "MTS",
  "UCS",
  "21F",
  "PDC",
  "MNR",
  "PAN-BOL",
];

const axios = require("axios").default;

const apiURL = "https://testing06.com:8080";

export default {
  components: {
    VeBar,
  },

  data: function() {
    return {
      sucre: {
        results: [],
        chartData: {
          columns: ["partido", "Votos"],
          rows: [],
        },
        chartExtend: {
          series(item) {
            item[0].data = item[0].data.map((v, index) => ({
              value: v,
              itemStyle: { color: COLOR_LIST[index] },
            }));
            return item;
          },
        },
        precincts: [],
        precinctValue: null,
        showPrecinctResult: false,
        precinctResult: [],
        districts: ["1", "2", "3", "4", "5"],
        districtValue: "",
        showDistrictResult: false,
        districtResult: [],
      },
    };
  },

  mounted() {
    axios
      .get(apiURL + "/api/v1/sw1/results-president-city/sucre")
      .then((response) => {
        var data = response.data;
        this.sucre.results = data;
        let n = 9;
        while (n > 0) {
          let partido = {
            partido: partidos[n - 1],
            Votos: data[n - 1].votos,
          };
          this.sucre.chartData.rows.push(partido);
          n--;
        }
      })
      .catch((error) => {
        console.log(error);
      });

    axios
      .get(apiURL + "/api/v1/sw1/recincts/sucre")
      .then((response) => {
        this.sucre.precincts = response.data;
      })
      .catch((error) => {
        console.log(error);
      });
  },

  methods: {
    customFilter(item, queryText, itemText) {
      console.log(itemText);
      const textOne = item.name.toLowerCase();
      const textTwo = item.abbr.toLowerCase();
      const searchText = queryText.toLowerCase();

      return (
        textOne.indexOf(searchText) > -1 || textTwo.indexOf(searchText) > -1
      );
    },

    getReportsucre: function() {
      axios
        .post(
          apiURL + "/api/v1/sw1/generate-report-president-city/sucre",
          this.sucre.results,
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

    getReportRecinctsucre: function() {
      console.log("Generando reporteeee");
    },

    getReportsucreDistrict: function() {
      axios
        .post(
          apiURL +
            "/api/v1/sw1/generate-report-president-district/" +
            this.sucre.districtValue,
          this.sucre.districtResult,
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

    getPrecinctResultsucre: function() {
      axios
        .post(apiURL + "/api/v1/sw1/result-from-precinct", {
          name: this.sucre.precinctValue,
        })
        .then((response) => {
          this.sucre.precinctResult = response.data;
          this.sucre.showPrecinctResult = true;
        })
        .catch((error) => {
          console.log(error);
        });
    },

    getDistrictResultsucre: function() {
      axios
        .post(apiURL + "/api/v1/sw1/result-from-district", {
          circunscripcion: this.sucre.districtValue,
        })
        .then((response) => {
          this.sucre.districtResult = response.data;
          this.sucre.showDistrictResult = true;
        })
        .catch((error) => {
          console.log(error);
        });
    },
  },
};
</script>