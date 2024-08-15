 Cesium.Ion.defaultAccessToken = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiIwNDUxMjM2MS00ODYwLTRjYjEtODhmMy0yZGE3NGVlMjUyOTkiLCJpZCI6MjExMzY3LCJpYXQiOjE3MTQwNzI0MzN9.54sPRmtK-snUlZx3mB3PXADPHVwc1X43K0ybFCvIHhA';

const viewer = new Cesium.Viewer('cesium', {
  terrain: Cesium.Terrain.fromWorldTerrain(),
  shadowMap: false,
  fullscreenButton: false,
  navigationHelpButton: false,
  vrButton: false,
  sceneModePicker: false,
  geocoder: true,
  infobox: false,
  selectionIndicator: false,
  timeline: false,
  projectionPicker: false,
  clockViewModel: null,
  animation: false,
  // https://cesium.com/blog/2018/01/24/cesium-scene-rendering-performance/#enabling-request-render-mode
  requestRenderMode: true,
});

//destroy both the skybox and the sun
viewer.scene.skyBox.destroy()
viewer.scene.skyBox = undefined
viewer.scene.sun.destroy();
viewer.scene.sun = undefined;
viewer.scene.moon.destroy();
viewer.scene.moon = undefined;
