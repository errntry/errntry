<template>
  <div class="bg-white h-full w-full text-blue-800 pt-8">
    <div class="font-bold text-3xl uppercase text-blue-700 text-center">
      {{ auth }}
    </div>
    <div
      class="w-full flex items-center justify-center mt-5 px-10 pd:mx-24 lg:px-40"
    >
      <form
        id="loginForm"
        class="block w-full lg:w-2/3 text-xl"
        action="#"
        method="post"
      >
        <div class="w-full">
          <label for="name" class="block text-blue-700">Name</label>
          <input
            id="name"
            class="block w-full px-6 py-1 transition-all duration-400 ease-in border rounded border-blue-700 outline-none focus:bg-blue-100"
            type="text"
            placeholder="Name"
            name="name"
            v-model="name"
            autocomplete="off"
            required
          />
        </div>
        <div class="mt-8 w-full">
          <label for="username" class="block text-blue-700">Username</label>
          <input
            id="username"
            class="block w-full px-6 py-1 transition-all duration-400 ease-in border rounded border-blue-700 outline-none focus:bg-blue-100"
            type="text"
            placeholder="Username"
            name="username"
            v-model="username"
            autocomplete="off"
            required
          />
        </div>
        <div class="mt-8 w-full">
          <label for="email" class="block text-blue-700">Email</label>
          <input
            id="email"
            class="block w-full px-6 py-1 transition-all duration-400 ease-in border rounded border-blue-700 outline-none focus:bg-blue-100"
            type="text"
            placeholder="Email"
            name="email"
            v-model="email"
            autocomplete="off"
            required
          />
        </div>
        <div class="mt-8 w-full">
          <label for="password" class="block text-blue-700">Password</label>
          <input
            id="password"
            class="block w-full px-6 py-1 transition-all duration-400 ease-in border rounded border-blue-700 outline-none focus:bg-blue-100"
            type="password"
            placeholder="Password"
            name="password"
            v-model="password"
            required
          />
        </div>
        <button
          type="submit"
          class="border mt-5 py-1 px-6 rounded bg-blue-700 text-white hover:bg-white hover:border-blue-700 hover:text-blue-700 hover:shadow"
        >
          {{ auth }}
        </button>
        <div class="text-center">
          {{ auth }} with Github or GitLab below or
          <router-link to="/login" class="underline">Login here</router-link>
          if you already have an {{ title }} account.
        </div>
        <button
          type="submit"
          class="border w-full flex items-center justify-center mt-5 py-2 px-6 rounded bg-blue-700 text-white hover:bg-white hover:border-blue-700 hover:text-blue-700 hover:shadow"
          disabled
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-5 w-auto fill-none stroke-current"
            viewBox="0 0 24 24"
            stroke-width="1.5"
            fill="none"
            stroke-linecap="round"
            stroke-linejoin="round"
          >
            <path stroke="none" d="M0 0h24v24H0z" fill="none" />
            <path
              d="M9 19c-4.3 1.4 -4.3 -2.5 -6 -3m12 5v-3.5c0 -1 .1 -1.4 -.5 -2c2.8 -.3 5.5 -1.4 5.5 -6a4.6 4.6 0 0 0 -1.3 -3.2a4.2 4.2 0 0 0 -.1 -3.2s-1.1 -.3 -3.5 1.3a12.3 12.3 0 0 0 -6.2 0c-2.4 -1.6 -3.5 -1.3 -3.5 -1.3a4.2 4.2 0 0 0 -.1 3.2a4.6 4.6 0 0 0 -1.3 3.2c0 4.6 2.7 5.7 5.5 6c-.6 .6 -.6 1.2 -.5 2v3.5"
            />
          </svg>
          Github
        </button>
        <button
          type="submit"
          class="border w-full flex items-center justify-center mt-5 py-2 px-6 rounded bg-blue-700 text-white hover:bg-white hover:border-blue-700 hover:text-blue-700 hover:shadow"
          disabled
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-5 w-auto fill-none stroke-current"
            viewBox="0 0 24 24"
            stroke-width="1.5"
            fill="none"
            stroke-linecap="round"
            stroke-linejoin="round"
          >
            <path stroke="none" d="M0 0h24v24H0z" fill="none" />
            <path d="M21 14l-9 7l-9 -7l3 -11l3 7h6l3 -7z" />
          </svg>
          Gitlab
        </button>
      </form>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";

export default defineComponent({
  name: "Signup",
  props: {
    title: {
      type: String,
      require: true
    }
  },
  setup() {
    const auth = ref(`Signup`);
    const name = ref(``);
    const username = ref(``);
    const email = ref(``);
    const password = ref(``);
    return {
      auth,
      name,
      username,
      email,
      password
    };
  }
});
</script>
