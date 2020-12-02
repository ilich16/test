<template>
  <div>
    <v-card class="ma-4">
      <v-card-title class="headline font-weight-bold justify-center">
        ACTAS POR DEPARTAMENTO
      </v-card-title>
      <div class="py-6">
        <ve-pie :data="chartData"></ve-pie>
      </div>
    </v-card>
    <v-card class="mx-auto" max-width="1000">
      <v-simple-table>
        <template v-slot:default>
          <thead>
            <tr>
              <th class="text-left">Ciudad</th>
              <th class="text-left">Cantidad de actas</th>
              <th class="text-left">%</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="item in dataTable" :key="item.ciudad">
              <td>{{ item.ciudad }}</td>
              <td>{{ item.actas }}</td>
              <td>{{ item.porcentaje }}</td>
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
      </v-card-actions>
    </v-card>
  </div>
</template>

<script>
import VePie from "v-charts/lib/pie.common";
import fileDownload from "js-file-download";

const axios = require("axios").default;
const apiURL = "https://testing06.com:8080";

export default {
  components: {
    VePie,
  },
  data: function() {
    return {
      dataTable: [],
      chartData: {
        columns: ["ciudad", "actas"],
        rows: [],
      },
    };
  },

  mounted() {
    axios.get(apiURL + "/api/v1/sw1/total-proceedings").then((response) => {
      this.dataTable = response.data;
      let chartData = [];
      for (let i = 0; i <= 8; i++) {
        chartData.push(response.data[i]);
      }
      this.$set(this.chartData, "rows", chartData);
    });
  },

  methods: {
    generateReport: function() {
      axios
        .post(
          apiURL +
            "/api/v1/sw1/generate-report-proceeding-city",
          this.dataTable,
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
        })
        .catch((error) => {
          console.log(error);
        });
    },
  }
};
</script>
