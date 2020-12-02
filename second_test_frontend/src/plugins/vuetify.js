import Vue from "vue";
import Vuetify from "vuetify/lib";

Vue.use(Vuetify);

import es from 'vuetify/es5/locale/es'

export default new Vuetify({
  lang: {
    locales: { es },
    current: 'es',
  }
});
