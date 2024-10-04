import { createRouter, createWebHistory } from "vue-router";


const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [

    {
      path: "/",
      name: "soal1",
      component: () => import("../views/Soal1View.vue"),
    },
    {
      path: "/soal1",
      name: "soal1",
      component: () => import("../views/Soal1View.vue"),
    },
    {
      path: "/soal2",
      name: "soal2",
      component: () => import("../views/Soal2View.vue"),
    },
  ],
});

export default router;
