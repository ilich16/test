<template>
  <div class="home">
    <v-app-bar color="deep-purple accent-4" dark>
      <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>

      <v-toolbar-title>Resultados - Diputados</v-toolbar-title>
    </v-app-bar>

    <v-navigation-drawer v-model="drawer" absolute bottom temporary>
      <v-list-item>
        <v-list-item-content>
          <v-list-item-title class="title">
            Diego Ilich
          </v-list-item-title>
          <v-list-item-subtitle>
            diego.severiche@gmail.com
          </v-list-item-subtitle>
        </v-list-item-content>
      </v-list-item>

      <v-divider></v-divider>

      <v-list nav dense>

        <v-list-item link to="resultados-presidenciales">
          <v-list-item-title>Resultados - Presidenciales</v-list-item-title>
        </v-list-item>

        <v-list-item link to="resultados-diputados">
          <v-list-item-title>Resultados - Diputados</v-list-item-title>
        </v-list-item>

        <v-list-item link to="actas">
          <v-list-item-title>Actas</v-list-item-title>
        </v-list-item>
      </v-list>
        <div class="pa-2">
          <v-btn block @click.stop="logoutDialog = true">Cerrar sesión</v-btn>
        </div>
    </v-navigation-drawer>

    <div class="pa-4">
      <v-card class="mx-auto" max-width="600">
        <v-card-text>
          <p class="display-1 text--primary">
            Buscar circunscripción
          </p>
          <div class="text--primary">
            ¿Estás buscando una circunscripción en específico? Busca en la
            siguiente lista de circunscripciones disponibles y luego presiona en
            el botón para cargar sus datos correspondientes.
          </div>
        </v-card-text>
        <v-autocomplete
          v-model="bolivia.districtValue"
          class="mx-4"
          :items="bolivia.districts"
          color="deep-purple accent-4"
          label="Circunscripción"
        ></v-autocomplete>
        <v-card-actions>
          <v-btn
            text
            color="deep-purple accent-4"
            @click="getDistrictResultGeneral"
          >
            Cargar datos
          </v-btn>
        </v-card-actions>
      </v-card>
    </div>
    <v-card flat v-if="bolivia.showDistrictResult">
      <div class="pa-2">
        <v-card class="ma-auto" max-width="400">
          <v-card-title class="headline font-weight-bold justify-center">
            {{ bolivia.title }}
          </v-card-title>
        </v-card>
      </div>
      <v-card class="ma-4">
        <ve-bar
          class="px-4"
          :data="bolivia.chartData"
          :extend="bolivia.chartExtend"
          :legend-visible="false"
        ></ve-bar>
      </v-card>
      <v-card class="ma-4">
        <v-simple-table>
          <template v-slot:default>
            <thead>
              <tr>
                <th class="text-left">Partido político</th>
                <th class="text-left">Votos</th>
                <th class="text-left">%</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="item in bolivia.districtResult" :key="item.partido">
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
            @click="getReportMembers"
          >
            Generar reporte
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-card>
    <v-dialog v-model="logoutDialog" max-width="290">
      <v-card>
        <v-card-title class="headline">Cerrar sesión</v-card-title>

        <v-card-text>
          ¿Estás seguro?
        </v-card-text>

        <v-card-actions>
          <v-spacer></v-spacer>

          <v-btn color="deep-purple accent-4" text @click="logoutDialog = false">
            Cancelar
          </v-btn>

          <v-btn color="deep-purple accent-4" text @click="logout">
            Aceptar
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>
// @ is an alias to /src
// import HelloWorld from "@/components/HelloWorld.vue";
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
  name: "Home",
  components: {
    VeBar,
  },
  data: function() {
    return {
      drawer: false,
      group: null,
      logoutDialog: false,
      bolivia: {
        title: "",
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
        districts: [],
        districtValue: null,
        showDistrictResult: false,
      },
      santaCruz: {
        title: "",
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
        districts: [
          "44",
          "45",
          "46",
          "47",
          "48",
          "49",
          "50",
          "51",
          "52",
          "53",
          "54",
          "55",
          "56",
          "57",
        ],
        districtValue: null,
        showDistrictResult: false,
      },
    };
  },

  mounted() {
    let districts = [];

    for (let i = 1; i <= 63; i++) {
      districts.push(i.toString());
    }

    this.bolivia.districts = districts;

    // Requests for 'Santa Cruz'
    axios
      .get(apiURL + "/api/v1/sw1/recincts/santa-cruz")
      .then((response) => {
        this.$set(this.santaCruz, "precincts", response.data);
      })
      .catch((error) => {
        console.log(error);
      });
  },

  methods: {
    // Bolivia
    getDistrictResultGeneral: function() {
      axios
        .post(apiURL + "/api/v1/sw1/member-result-from-district", {
          circunscripcion: this.bolivia.districtValue,
        })
        .then((response) => {
          var data = response.data;
          this.$set(this.bolivia, "districtResult", data);
          let chartData = [];
          let n = 9;
          while (n > 0) {
            let partido = {
              partido: partidos[n - 1],
              Votos: data[n - 1].votos,
            };
            chartData.push(partido);
            n--;
          }
          this.$set(this.bolivia.chartData, "rows", chartData);
          this.$set(
            this.bolivia,
            "title",
            "Circunscripción " + this.bolivia.districtValue
          );
          this.bolivia.showDistrictResult = true;
        })
        .catch((error) => {
          console.log(error);
        });
    },

    // Santa Cruz
    getDistrictResultSantaCruz: function() {
      axios
        .post(apiURL + "/api/v1/sw1/member-result-from-district", {
          circunscripcion: this.santaCruz.districtValue,
        })
        .then((response) => {
          var data = response.data;
          this.$set(this.santaCruz, "districtResult", data);
          let chartData = [];
          let n = 9;
          while (n > 0) {
            let partido = {
              partido: partidos[n - 1],
              Votos: data[n - 1].votos,
            };
            chartData.push(partido);
            n--;
          }
          this.$set(this.santaCruz.chartData, "rows", chartData);
          this.$set(
            this.santaCruz,
            "title",
            "Circunscripción " + this.santaCruz.districtValue
          );
          this.santaCruz.showDistrictResult = true;
        })
        .catch((error) => {
          console.log(error);
        });
    },

    getPrecinctResultSantaCruz: function() {
      axios.post(apiURL + "/api/v1/sw/member-result-from-precinct");
    },

    getReportMembers: function() {
      axios
        .post(
          apiURL +
            "/api/v1/sw1/generate-report-member-district/" +
            this.bolivia.districtValue,
          this.bolivia.districtResult,
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

    logout: function() {
      this.$router.push('/login')
    }
  },

  watch: {
    group: function() {
      this.drawer = false;
    },

    bolivia: function() {
      console.log("Ha cambiadodooooo");
    },
  },
};
</script>
