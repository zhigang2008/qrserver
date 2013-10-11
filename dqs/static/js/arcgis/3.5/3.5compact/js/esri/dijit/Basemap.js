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
define("esri/dijit/Basemap",["dojo/_base/declare","dojo/_base/array","dojo/_base/lang","dojo/has","esri/kernel","esri/request","esri/dijit/BasemapLayer"],function(_1,_2,_3,_4,_5,_6,_7){var BM=_1(null,{declaredClass:"esri.dijit.Basemap",id:null,title:"",thumbnailUrl:null,layers:null,itemId:null,basemapGallery:null,constructor:function(_8,_9){_8=_8||{};if(!_8.layers&&!_8.itemId){console.error("esri.dijit.Basemap: unable to find the 'layers' property in parameters");}this.id=_8.id;this.itemId=_8.itemId;this.layers=_8.layers;this.title=_8.title||"";this.thumbnailUrl=_8.thumbnailUrl;this.basemapGallery=_9;},getLayers:function(){if(this.layers){return this.layers;}else{if(this.itemId){var _a=_5.dijit._arcgisUrl+"/content/items/"+this.itemId+"/data";var _b={};_b.f="json";var _c=_6({url:_a,content:_b,callbackParamName:"callback",error:_3.hitch(this,function(_d,_e){var _f="esri.dijit.Basemap: could not access basemap item.";if(this.basemapGallery){this.basemapGallery.onError(_f);}else{console.error(_f);}})});_c.addCallback(_3.hitch(this,function(_10,_11){if(_10.baseMap){this.layers=[];_2.forEach(_10.baseMap.baseMapLayers,function(_12){var _13={};if(_12.url){_13.url=_12.url;}if(_12.type){_13.type=_12.type;}if(_12.isReference){_13.isReference=_12.isReference;}if(_12.displayLevels){_13.displayLevels=_12.displayLevels;}if(_12.visibleLayers){_13.visibleLayers=_12.visibleLayers;}if(_12.bandIds){_13.bandIds=_12.bandIds;}this.layers.push(new _7(_13));},this);return this.layers;}else{var msg="esri.dijit.Basemap: could not access basemap item.";if(this.basemapGallery){this.basemapGallery.onError(msg);}else{console.error(msg);}return [];}}));return _c;}}}});if(_4("extend-esri")){_3.setObject("dijit.Basemap",BM,_5);}return BM;});