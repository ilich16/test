<template>
  <div>
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
      <template v-slot:append>
        <div class="pa-2">
          <v-btn block @click.stop="logoutDialog = true">Cerrar sesión</v-btn>
        </div>
      </template>
    </v-navigation-drawer>

    <v-dialog v-model="logoutDialog" max-width="290">
      <v-card>
        <v-card-title class="headline">Cerrar sesión</v-card-title>

        <v-card-text>
          ¿Estás seguro?
        </v-card-text>

        <v-card-actions>
          <v-spacer></v-spacer>

          <v-btn color="green darken-1" text @click="logoutDialog = false">
            Cancelar
          </v-btn>

          <v-btn color="green darken-1" text @click="logoutDialog = false">
            Aceptar
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>


export default {
  data: function() {
    return {
      drawer: false,
      group: null,
      logoutDialog: false,
    };
  },

  watch: {
    group: function() {
      this.drawer = false;
    },
  },
};
</script>