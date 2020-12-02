<template>
  <div class="home">
    
    <navigation></navigation>

    <proceeding-city></proceeding-city>

    <proceeding-information></proceeding-information>

  </div>
</template>

<script>

import Navigation from "../components/navigation/proceedings/Navigation";
import ProceedingCity from "../components/proceedings/ProceedingCity.vue";
import ProceedingInformation from "../components/proceedings/ProceedingInformation.vue";

export default {
  name: "Home",
  components: {
    Navigation,
    ProceedingCity,
    ProceedingInformation
  },
  data: function() {
    return {
    };
  },
};
</script>
