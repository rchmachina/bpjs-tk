import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: HomeView,
    },
    {
      path: "/soal1",
      name: "tableView",
      component: () => import("../views/TableView.vue"),
    },
    {
      path: "/soal2",
      name: "calculator",
      component: () => import("../views/CalculatorView.vue"),
    },
  ],
});

export default router;
