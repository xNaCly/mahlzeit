// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  app: {
    buildAssetsDir: "/nuxt/",
    head: {
      title: "mahlzeit",
    },
  },
  devtools: { enabled: true },
  modules: ["@nuxt/ui"],
  compatibilityDate: "2025-01-03",
  css: ["~/assets/main.css"],
  runtimeConfig: {
    public: {
      server:
        process.env.NODE_ENV === "production"
          ? "/api/"
          : "http://localhost:8080/api/",
    },
  },
  colorMode: {
    preference: "light",
  },
});
