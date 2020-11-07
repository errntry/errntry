<template>
  <div class="w-full h-full bg-blue-700">
    <div class="absolute flex items-center justify-center w-full h-full px-5">
      <div class="text-center text-white">
        <h1 class="text-6xl font-bold font-oleo">
          {{ props.title }}
        </h1>
        <p class="text-lg ">
          {{ props.motto }}
        </p>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "Home",
  components: {},
  props: {
    title: String,
    motto: String
  },
  setup(props: object) {
    return {
      props
    };
  }
});
</script>
