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
define("esri/layers/JoinDataSource",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel","esri/lang","esri/layers/DataSource","esri/layers/LayerMapSource","esri/layers/TableDataSource","esri/layers/QueryDataSource","esri/layers/RasterDataSource"],function(_1,_2,_3,_4,_5,_6,_7,_8,_9,_a){var _b=_1(_6,{declaredClass:"esri.layers.JoinDataSource",constructor:function(_c){if(_c){if(_c.leftTableSource){this.leftTableSource=this._createLayerSource(_c.leftTableSource);}if(_c.rightTableSource){this.rightTableSource=this._createLayerSource(_c.rightTableSource);}}},_createLayerSource:function(_d){var _e;if(_d.type==="mapLayer"){_e=new _7(_d);}else{_e={type:"dataLayer"};var _f;switch(_d.dataSource.type){case "table":_f=new _8(_d.dataSource);break;case "queryTable":_f=new _9(_d.dataSource);break;case "joinTable":_f=new _b(_d.dataSource);break;case "raster":_f=new _a(_d.dataSource);break;default:_f=_d.dataSource;}_e.dataSource=_f;_e.toJson=function(){var _10={type:"dataLayer",dataSource:_f.toJson()};return _5.fixJson(_10);};}return _e;},toJson:function(){var _11={type:"joinTable",leftTableSource:this.leftTableSource&&this.leftTableSource.toJson(),rightTableSource:this.rightTableSource&&this.rightTableSource.toJson(),leftTableKey:this.leftTableKey,rightTableKey:this.rightTableKey};var _12;if(this.joinType.toLowerCase()==="left-outer-join"){_12="esriLeftOuterJoin";}else{if(this.joinType.toLowerCase()==="left-inner-join"){_12="esriLeftInnerJoin";}else{_12=this.joinType;}}_11.joinType=_12;return _5.fixJson(_11);}});if(_3("extend-esri")){_2.setObject("layers.JoinDataSource",_b,_4);}return _b;});