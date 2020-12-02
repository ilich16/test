<template>
  <div class="home">
    <v-app-bar color="deep-purple accent-4" dark>
      <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>

      <v-toolbar-title>Resultados - Presidenciales</v-toolbar-title>
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

    <v-tabs background-color="deep-purple accent-4" center-active dark grow>
      <v-tab>General</v-tab>
      <v-tab>Santa Cruz</v-tab>
      <v-tab>La Paz</v-tab>
      <v-tab>Cochabamba</v-tab>
      <v-tab>Sucre</v-tab>
      <v-tab>Tarija</v-tab>
      <v-tab>Pando</v-tab>
      <v-tab>Beni</v-tab>
      <v-tab>Oruro</v-tab>
      <v-tab>Potosí</v-tab>
      <!-- Item for 'General' -->
      <v-tab-item>
        <general></general>
      </v-tab-item>
      <!-- Item for 'Santa Cruz' -->
      <v-tab-item>
        <santa-cruz></santa-cruz>
      </v-tab-item>
      <!-- Item for 'La Paz' -->
      <v-tab-item>
        <la-paz></la-paz>
      </v-tab-item>
      <!-- Item for 'Cochabamba' -->
      <v-tab-item>
        <cochabamba></cochabamba>
      </v-tab-item>
      <!-- Item for 'Sucre' -->
      <v-tab-item>
        <sucre></sucre>
      </v-tab-item>
      <!-- Item for 'Tarija' -->
      <v-tab-item>
        <tarija></tarija>
      </v-tab-item>
      <!-- Item for 'Pando' -->
      <v-tab-item>
        <pando></pando>
      </v-tab-item>
      <!-- Item for 'Beni' -->
      <v-tab-item>
        <beni></beni>
      </v-tab-item>
      <!-- Item for 'Oruro' -->
      <v-tab-item>
        <oruro></oruro>
      </v-tab-item>
      <!-- Item for 'Potosí' -->
      <v-tab-item>
        <potosi></potosi>
      </v-tab-item>
    </v-tabs>
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
import General from "../components/results/president/general/General.vue";
import SantaCruz from "../components/results/president/santacruz/SantaCruz.vue";
import LaPaz from "../components/results/president/lapaz/LaPaz.vue";
import Cochabamba from "../components/results/president/cochabamba/Cochabamba.vue";
import Sucre from "../components/results/president/sucre/Sucre.vue";
import Tarija from "../components/results/president/tarija/Tarija.vue";
import Pando from "../components/results/president/pando/Pando.vue";
import Beni from "../components/results/president/beni/Beni.vue";
import Oruro from "../components/results/president/oruro/Oruro.vue";
import Potosi from "../components/results/president/potosi/Potosi.vue";

export default {
  components: {
    General,
    SantaCruz,
    LaPaz,
    Cochabamba,
    Sucre,
    Tarija,
    Pando,
    Beni,
    Oruro,
    Potosi,
  },
  data: function() {
    return {
      drawer: false,
      group: null,
      logoutDialog: false,
    };
  },

  methods: {
    logout: function() {
      this.$router.push('/login')
    }
  }

};
</script>
