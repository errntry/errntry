<template>
  <div class="w-screen h-screen antialiased font-mulish">
    <router-view :title="title" :motto="motto" :key="$route.fullPath" />
    <router-link :title="title" :motto="motto" to="/"></router-link>
  </div>
</template>
<script lang="ts">
import { ref } from "vue";

export default {
  setup() {
    const title = ref(`errntry`);
    const motto = ref(
      `An open source cross-platform application monitoring with a focus on error reporting`
    );
    return {
      title,
      motto
    };
  }
};
</script>
