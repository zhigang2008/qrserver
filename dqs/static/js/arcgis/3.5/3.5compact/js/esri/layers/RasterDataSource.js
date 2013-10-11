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
define("esri/layers/RasterDataSource",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel","esri/lang","esri/layers/DataSource"],function(_1,_2,_3,_4,_5,_6){var _7=_1(_6,{declaredClass:"esri.layers.RasterDataSource",toJson:function(){var _8={type:"raster",workspaceId:this.workspaceId,dataSourceName:this.dataSourceName};return _5.fixJson(_8);}});if(_3("extend-esri")){_2.setObject("layers.RasterDataSource",_7,_4);}return _7;});