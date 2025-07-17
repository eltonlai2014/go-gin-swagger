window.onload = () => {
  window.ui = SwaggerUIBundle({
    url: "swagger.yaml",
    dom_id: '#swagger-ui',
    displayExplorer: false,  // ← 加這一行
    presets: [
      SwaggerUIBundle.presets.apis,
      SwaggerUIStandalonePreset
    ],
    layout: "BaseLayout"
  });
};