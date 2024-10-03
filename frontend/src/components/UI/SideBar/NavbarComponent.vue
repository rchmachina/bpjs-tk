<template>
  <nav
    class="bg-teal-500 text-center text-white float-left fixed md:static min-h-screen z-10 top-0 left-0 bottom-0 p-2 flex flex-col transition-all duration-700 ease"
    :class="sizeNavbar"
  >
    <i
      class="bx bxs-chevrons-left bx-sm cursor-pointer absolute p-2 text-green-200 transition-all duration-500 ease-linear rotate-0"
      :class="rotateIcon"
      @click="collapsedNavbar"
    ></i>

    <h1 class="text-3xl pt-10">
      <span v-if="stateSideBar">
        <router-link to="/">
          <img
            class="bg-green-300 rounded-full hover:bg-green-400 cursor-pointer"
            src=""
          />
        </router-link>
      </span>
      <span v-else>BPJS TK</span>
    </h1>

    <section class="flex flex-col gap-5 mt-10">
      <SideBarLink to="/" text="Home" icon="bx-home" />
      <SideBarLink to="/soal1" text="Home" icon="bx-home" />
      <SideBarLink to="/soal2" text="table" icon="bx-table" />
    </section>
  </nav>
</template>

<script>
import { computed } from "@vue/runtime-core";
import SideBarLink from "./SideBarLink.vue";
import { useStore } from "vuex";

export default {
  setup() {
    const store = useStore();
    const stateSideBar = computed(() => store.getters["sideBarCollapsed"]);

    const collapsedNavbar = () => store.commit("changeSideBar");
    const sizeNavbar = computed(() => (stateSideBar.value ? "w-14" : "w-80"));
    const rotateIcon = computed(() =>
      stateSideBar.value ? "rotate-180" : "rotate-0"
    );

    return {
      sizeNavbar,
      rotateIcon,
      stateSideBar,
      collapsedNavbar,
    };
  },
  components: { SideBarLink },
};
</script>
