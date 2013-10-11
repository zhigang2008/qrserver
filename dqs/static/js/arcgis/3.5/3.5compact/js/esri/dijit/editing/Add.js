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
define("esri/dijit/editing/Add",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel","esri/OperationBase"],function(_1,_2,_3,_4,_5){var _6=_1(_5,{declaredClass:"esri.dijit.editing.Add",type:"edit",label:"Add Features",constructor:function(_7){_7=_7||{};if(!_7.featureLayer){console.error("In constructor of 'esri.dijit.editing.Add', featureLayer is not provided");return;}this._featureLayer=_7.featureLayer;if(!_7.addedGraphics){console.error("In constructor of 'esri.dijit.editing.Add', no graphics provided");return;}this._addedGraphics=_7.addedGraphics;},performUndo:function(){this._featureLayer.applyEdits(null,null,this._addedGraphics);},performRedo:function(){this._featureLayer.applyEdits(this._addedGraphics,null,null);}});if(_3("extend-esri")){_2.setObject("dijit.editing.Add",_6,_4);}return _6;});