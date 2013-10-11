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
define("esri/layers/LayerDataSource",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel","esri/lang","esri/layers/LayerSource","esri/layers/TableDataSource","esri/layers/QueryDataSource","esri/layers/JoinDataSource","esri/layers/RasterDataSource"],function(_1,_2,_3,_4,_5,_6,_7,_8,_9,_a){var _b=_1(_6,{declaredClass:"esri.layers.LayerDataSource",type:"dataLayer",constructor:function(_c){if(_c&&_c.dataSource){var _d;switch(_c.dataSource.type){case "table":_d=new _7(_c.dataSource);break;case "queryTable":_d=new _8(_c.dataSource);break;case "joinTable":_d=new _9(_c.dataSource);break;case "raster":_d=new _a(_c.dataSource);break;default:_d=_c.dataSource;}this.dataSource=_d;}},toJson:function(){var _e={type:"dataLayer",dataSource:this.dataSource&&this.dataSource.toJson()};return _5.fixJson(_e);}});if(_3("extend-esri")){_2.setObject("layers.LayerDataSource",_b,_4);}return _b;});