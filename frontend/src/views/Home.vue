<template>
  <div class="w-full h-full bg-blue-700 text-blue-100">
    <!--  login and signup routes  -->
    <span class="absolute top-0 right-0 flex m-3">
      <router-link
        to="/signup"
        class="uppercase px-3 py-1 text-current mx-1 rounded-sm hover:bg-gray-100 z-10 hover:text-blue-700 md:text-blue-700 md:hover:bg-blue-700 md:hover:text-gray-100 font-bold transition transition-colors duration-200"
      >
        signup
      </router-link>
      <router-link
        to="/login"
        class="uppercase px-5 py-1 text-current mx-1 rounded-sm hover:bg-gray-100 z-10 hover:text-blue-700 md:text-blue-700 md:hover:bg-blue-700 md:hover:text-gray-100 font-bold transition transition-colors duration-200"
      >
        login
      </router-link>
    </span>

    <!-- main body -->
    <div class="flex items-center justify-center w-full h-full bg-white">
      <div
        class="relative flex items-center justify-center w-full md:w-1/2 h-full bg-blue-700 md:clip-object"
      >
        <div class="text-center text-current px-10 lg:px-20">
          <h1 class="text-6xl font-bold font-oleo">
            {{ title }}
          </h1>
          <span class="text-lg "> {{ motto }}. </span>
        </div>
        <!-- copyright -->
        <span class="absolute bottom-0 inset-x-0 text-center mb-10 text-xl">
          <span class="font-oleo">{{ title }}</span> &copy;
          {{ new Date().getFullYear() }}</span
        >
      </div>
      <div
        class="hidden md:h-full md:w-1/2 md:flex md:items-center md:justify-center bg-white"
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
