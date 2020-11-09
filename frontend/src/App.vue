<template>
  <div class="w-screen h-screen antialiased font-mulish">
    <router-view :title="title" :motto="motto" :key="$route.fullPath" />
    <router-link to="/"></router-link>
  </div>
</template>
<script lang="ts">
import { defineComponent, ref } from "vue";

export default defineComponent({
  name: "App",
  components: {},
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
});
</script>
