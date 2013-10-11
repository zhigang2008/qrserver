/*
 COPYRIGHT 2009 ESRI

 TRADE SECRETS: ESRI PROPRIETARY AND CONFIDENTIAL
 Unpublished material - all rights reserved under the
 Copyright Laws of the United States and applicable international
 laws, treaties, and conventions.

 For additional information, contact:
 Environmental Systems Research Institute, Inc.
 Attn: Contracts and Legal Services Department
 380 New York Street
 Redlands, California, 92373
 USA

 email: contracts@esri.com
 */
//>>built
define("esri/config",["dojo/_base/config","dojo/topic","dojo/has","esri/kernel"],function(_1,_2,_3,_4){var _5={defaults:{screenDPI:96,geometryService:null,kmlService:null,map:{width:400,height:400,layerNamePrefix:"layer",graphicsLayerNamePrefix:"graphicsLayer",slider:{left:"30px",top:"30px",width:null,height:"200px"},sliderLabel:{tick:5,labels:null,style:"width:2em; font-family:Verdana; font-size:75%;"},sliderChangeImmediate:true,zoomSymbol:{color:[0,0,0,64],outline:{color:[255,0,0,255],width:1.25,style:"esriSLSSolid"},style:"esriSFSSolid"},zoomDuration:500,zoomRate:25,panDuration:350,panRate:25,logoLink:"http://www.esri.com",basemaps:{"streets":{title:"Streets",baseMapLayers:[{url:"http://services.arcgisonline.com/ArcGIS/rest/services/World_Street_Map/MapServer"}]},"satellite":{title:"Satellite",baseMapLayers:[{url:"http://services.arcgisonline.com/ArcGIS/rest/services/World_Imagery/MapServer"}]},"hybrid":{title:"Imagery with Labels",baseMapLayers:[{url:"http://services.arcgisonline.com/ArcGIS/rest/services/World_Imagery/MapServer"},{url:"http://services.arcgisonline.com/ArcGIS/rest/services/Reference/World_Boundaries_and_Places/MapServer",isReference:true}]},"topo":{title:"Topographic",baseMapLayers:[{url:"http://services.arcgisonline.com/ArcGIS/rest/services/World_Topo_Map/MapServer"}]},"gray":{title:"Light Gray Canvas",baseMapLayers:[{url:"http://services.arcgisonline.com/ArcGIS/rest/services/Canvas/World_Light_Gray_Base/MapServer"},{url:"http://services.arcgisonline.com/ArcGIS/rest/services/Canvas/World_Light_Gray_Reference/MapServer",isReference:true}]},"oceans":{title:"Oceans",baseMapLayers:[{url:"http://services.arcgisonline.com/ArcGIS/rest/services/Ocean_Basemap/MapServer"}]},"national-geographic":{title:"National Geographic",baseMapLayers:[{url:"http://services.arcgisonline.com/ArcGIS/rest/services/NatGeo_World_Map/MapServer"}]},"osm":{title:"OpenStreetMap",baseMapLayers:[{type:"OpenStreetMap"}]}}},io:{errorHandler:function(_6,io){_2.publish("esri.Error",[_6]);},proxyUrl:null,alwaysUseProxy:false,corsEnabledServers:["www.arcgis.com","tiles.arcgis.com","services.arcgis.com","services1.arcgis.com","services2.arcgis.com","services3.arcgis.com","static.arcgis.com","utility.arcgisonline.com","geocode.arcgis.com","qaext.arcgis.com","tilesqa.arcgis.com","servicesqa.arcgis.com","servicesqa1.arcgis.com","servicesqa2.arcgis.com","servicesqa3.arcgis.com","geocodeqa.arcgis.com","dev.arcgis.com","devext.arcgis.com","tilesdevext.arcgis.com","servicesdev.arcgis.com","servicesdev1.arcgis.com","servicesdev2.arcgis.com","servicesdev3.arcgis.com","geocodedev.arcgis.com"],corsDetection:true,_processedCorsServers:{},proxyRules:[],postLength:2000,timeout:60000}}};if(_3("extend-esri")){_4.config=_5;}if(!_1.noGlobals){window.esriConfig=_5;}return _5;});