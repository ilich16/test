<template>
  <div>
      <v-card flat>
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
                  <th class="text-left">
                    Candidato a presidente y vicepresidente
                  </th>
                  <th class="text-left">Partido</th>
                  <th class="text-left">Votos</th>
                  <th class="text-left">%</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="item in bolivia.results" :key="item.candidato">
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
              @click="getReportGeneral"
            >
              Generar reporte
            </v-btn>
          </v-card-actions>
        </v-card>
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
      bolivia: {
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
      },
    };
  },

  mounted() {
    axios
      .get(apiURL + "/api/v1/sw1/results-president-country")
      .then((response) => {
        var data = response.data;
        this.bolivia.results = data;
        let n = 9;
        while (n > 0) {
          let partido = {
            partido: partidos[n - 1],
            Votos: data[n - 1].votos,
          };
          this.bolivia.chartData.rows.push(partido);
          n--;
        }
      })
      .catch((error) => {
        console.log(error);
      });
  },

  methods: {
    getReportGeneral: function() {
      axios
        .post(
          apiURL + "/api/v1/sw1/generate-report-president-country",
          this.bolivia.results,
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
  }
};
</script>
