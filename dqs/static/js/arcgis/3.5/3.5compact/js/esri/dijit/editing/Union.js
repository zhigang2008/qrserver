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
define("esri/dijit/editing/Union",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel","esri/OperationBase","esri/dijit/editing/Cut"],function(_1,_2,_3,_4,_5,_6){var _7=_1(_5,{declaredClass:"esri.dijit.editing.Union",type:"edit",label:"Union Features",constructor:function(_8){_8=_8||{};this._cut=new _6({featureLayer:_8.featureLayer,addedGraphics:_8.deletedGraphics,preUpdatedGraphics:_8.preUpdatedGraphics,postUpdatedGraphics:_8.postUpdatedGraphics});},performUndo:function(){this._cut.performRedo();},performRedo:function(){this._cut.performUndo();}});if(_3("extend-esri")){_2.setObject("dijit.editing.Union",_7,_4);}return _7;});