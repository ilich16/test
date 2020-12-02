import Vue from "vue";
import VueRouter from "vue-router";
const Login = () => import("../views/Login.vue")
const PresidentResults = () => import("../views/PresidentResults.vue")
const MemberResults = () => import("../views/MemberResults.vue")
const Proceedings = () => import("../views/Proceedings.vue")

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    redirect: "/login"
  },
  {
    path: "/login",
    name: "Login",
    component: Login
  },
  {
    path: "/resultados-presidenciales",
    name: "PresidentResults",
    component: PresidentResults
  },
  {
    path: "/resultados-diputados",
    name: "MemberResults",
    component: MemberResults
  },
  {
    path: "/actas",
    name: "Proceedings",
    component: Proceedings
  },
];

const router = new VueRouter({
  mode: 'history',
  routes
});

export default router;
