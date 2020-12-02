<template>
  <v-dialog v-model="dialog" persistent max-width="300">
    <v-card>
      <v-card-title>
        <span class="headline">{{ GetDialogTitle }}</span>
      </v-card-title>
      <v-card-text>
        {{ GetDialogMessage }}
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="deep-purple accent-4" text v-on:click="CloseDialog"
          >Aceptar</v-btn
        >
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
export default {
  name: "result-dialog",
  props: {
    resultDialog: {
      type: Boolean,
      default: false,
    },
    responseCode: {
      type: Number,
      default: 0,
    },
  },
  data: function() {
    return {
      code: 0,
      dialog: false,
    };
  },
  watch: {
    responseCode: function(val) {
      this.code = val
    },
    resultDialog: function(val) {
      this.dialog = val;
    },
  },

  computed: {

    GetDialogTitle: function() {
      if (this.code == 200) {
        return "Iniciar sesión"
      } else {
        return "Error"
      }
    },

    GetDialogMessage: function() {
      if (this.code == 200) {
        return "Inicio de sesión realizado correctamente."
      } else {
        return "Por favor verifique que los datos sean correctos."
      }
    }

  },
  methods: {
    CloseDialog: function() {
      if (this.responseCode == 200) {
        this.$emit('close-dialog', false)
        this.$router.push("/resultados-presidenciales");
      } else {
        this.$emit('close-dialog', false)
      }
    },
  },
};
</script>
