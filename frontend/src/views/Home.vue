<template>
  <div class="w-full h-full text-blue-100 bg-blue-700">
    <!--  login and signup routes  -->
    <span class="absolute top-0 right-0 flex m-3">
      <router-link
        to="/signup"
        class="z-10 px-3 py-1 mx-1 font-bold text-current uppercase transition transition-colors duration-200 rounded-sm hover:bg-gray-100 hover:text-blue-700 md:text-blue-700 md:hover:bg-blue-700 md:hover:text-gray-100"
      >
        signup
      </router-link>
      <router-link
        to="/login"
        class="z-10 px-5 py-1 mx-1 font-bold text-current uppercase transition transition-colors duration-200 rounded-sm hover:bg-gray-100 hover:text-blue-700 md:text-blue-700 md:hover:bg-blue-700 md:hover:text-gray-100"
      >
        login
      </router-link>
    </span>

    <!-- main body -->
    <div class="flex items-center justify-center w-full h-full bg-white">
      <div
        class="relative flex items-center justify-center w-full h-full bg-blue-700 md:w-1/2 md:clip-object"
      >
        <div class="px-10 text-center text-current lg:px-20">
          <h1 class="text-6xl font-bold font-oleo">
            {{ title }}
          </h1>
          <span class="text-lg "> {{ motto }}. </span>
        </div>
        <!-- copyright -->
        <span class="absolute inset-x-0 bottom-0 mb-10 text-xl text-center">
          <span class="font-oleo">{{ title }}</span> &copy;
          {{ new Date().getFullYear() }}</span
        >
      </div>
      <div
        class="hidden bg-white md:h-full md:w-1/2 md:flex md:items-center md:justify-center"
      >
        <img src="../assets/undraw_bug_fixing_oc7a.png" alt="bug-fix" />
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
    title: {
      type: String,
      required: true
    },
    motto: {
      type: String,
      required: true
    }
  },
  setup() {
    return {};
  }
});
</script>
