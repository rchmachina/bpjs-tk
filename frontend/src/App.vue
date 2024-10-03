<template class="container mx-auto">
    <header class="header">
    <h1>Website Title</h1>
  </header>
  <div>
    <NavbarComponent />
    <div class="bg-amber-700 h-screen">
      <router-view/>
    </div>
  </div>
  <Footer></Footer>
</template>
<script>
import NavbarComponent from "./components/UI/SideBar/NavbarComponent.vue";
import Footer from "./components/UI/Footer/Footer.vue";
export default {
  setup() {},
  components: { NavbarComponent,Footer },
};
</script>
