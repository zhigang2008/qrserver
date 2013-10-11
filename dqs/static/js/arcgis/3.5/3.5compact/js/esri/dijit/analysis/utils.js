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
define("esri/dijit/analysis/utils",["dojo/_base/declare","dojo/_base/lang","dojo/_base/array","dojo/_base/connect","dojo/_base/event","dojo/_base/json","dojo/dom-attr","dojo/has","dojo/i18n","dojo/i18n!esri/nls/jsapi","dojo/json","dojo/query","dijit/registry","esri/kernel","esri/dijit/analysis/HelpWindow"],function(_1,_2,_3,_4,_5,_6,_7,_8,_9,_a,_b,_c,_d,_e,_f){var _10={};_2.mixin(_10,{initHelpLinks:function(_11){if(!esri.dijit._helpDialog){esri.dijit._helpDialog=new _f();}if(!_11){return;}var _12=_d.byNode(_11);var _13=_12.get("helpFileName");_c("[esriHelpTopic]",_11).forEach(function(_14,_15,_16){if(_14){_4.connect(_14,"onclick",_2.hitch(this,function(e){_5.stop(e);esri.dijit._helpDialog.show(e,_7.get(_14,"esriHelpTopic"),_13);}));}},this);},constructAnalysisFeatColl:function(_17){var obj={};obj.featureCollection=_17.layerDefinition;for(props in obj.featureCollection){if(props==="objectIdField"){obj.featureCollection.objectIdFieldName=_2.clone(obj.featureCollection.objectIdField);delete obj.featureCollection.objectIdField;}}obj.featureCollection.features=_17.featureSet.features;return obj;},constructAnalysisInputLyrObj:function(_18){var obj={};if(_18.url){obj={url:_18.url};if(_18.getDefinitionExpression&&_18.getDefinitionExpression()){obj.filter=_18.getDefinitionExpression();}if(_18.credential){obj.serviceToken=_18.credential.token;}}else{if(!_18.url){obj=this.constructAnalysisFeatColl(_18.toJson());}}return obj;}});if(_8("extend-esri")){_2.setObject("dijit.analysis.utils",_10,_e);}return _10;});