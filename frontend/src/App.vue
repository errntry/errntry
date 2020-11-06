<template>
  <div class="">
    {{ title }}
  </div>
</template>

<script lang="ts">
import { ref } from "vue";

export default {
  setup() {
    const title = ref(`errntry`);

    return {
      title
    };
  }
};
</script>
