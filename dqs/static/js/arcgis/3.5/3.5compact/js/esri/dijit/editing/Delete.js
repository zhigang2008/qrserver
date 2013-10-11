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
define("esri/dijit/editing/Delete",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel","esri/OperationBase","esri/dijit/editing/Add"],function(_1,_2,_3,_4,_5,_6){var _7=_1(_5,{declaredClass:"esri.dijit.editing.Delete",type:"edit",label:"Delete Features",constructor:function(_8){_8=_8||{};this._add=new _6({featureLayer:_8.featureLayer,addedGraphics:_8.deletedGraphics});},performUndo:function(){this._add.performRedo();},performRedo:function(){this._add.performUndo();}});if(_3("extend-esri")){_2.setObject("dijit.editing.Delete",_7,_4);}return _7;});