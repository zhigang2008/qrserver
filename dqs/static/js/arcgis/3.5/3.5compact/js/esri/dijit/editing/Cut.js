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
define("esri/dijit/editing/Cut",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel","esri/geometry/jsonUtils","esri/OperationBase"],function(_1,_2,_3,_4,_5,_6){var _7=_1(_6,{declaredClass:"esri.dijit.editing.Cut",type:"edit",label:"Cut Features",constructor:function(_8){var i;_8=_8||{};if(!_8.featureLayer){console.error("In constructor of 'esri.dijit.editing.Cut', featureLayer not provided");return;}this._featureLayer=_8.featureLayer;if(!_8.addedGraphics){console.error("In constructor of 'esri.dijit.editing.Cut', addedGraphics for cut not provided");return;}this._addedGraphics=_8.addedGraphics;if(!_8.preUpdatedGraphics){console.error("In constructor of 'esri.dijit.editing.Cut', preUpdatedGraphics not provided");return;}this._preUpdatedGraphicsGeometries=[];this._preUpdatedGraphicsAttributes=[];for(i=0;i<_8.preUpdatedGraphics.length;i++){this._preUpdatedGraphicsGeometries.push(_8.preUpdatedGraphics[i].geometry.toJson());this._preUpdatedGraphicsAttributes.push(_8.preUpdatedGraphics[i].attributes);}if(!_8.postUpdatedGraphics){console.error("In constructor of 'esri.dijit.editing.Cut', postUpdatedGraphics not provided");return;}this._postUpdatedGraphics=_8.postUpdatedGraphics;this._postUpdatedGraphicsGeometries=[];this._postUpdatedGraphicsAttributes=[];for(i=0;i<_8.postUpdatedGraphics.length;i++){this._postUpdatedGraphicsGeometries.push(_8.postUpdatedGraphics[i].geometry.toJson());this._postUpdatedGraphicsAttributes.push(_2.clone(_8.postUpdatedGraphics[i].attributes));}},performUndo:function(){var i;for(i=0;i<this._postUpdatedGraphics.length;i++){this._postUpdatedGraphics[i].setGeometry(_5.fromJson(this._preUpdatedGraphicsGeometries[i]));this._postUpdatedGraphics[i].setAttributes(this._preUpdatedGraphicsAttributes[i]);}this._featureLayer.applyEdits(null,this._postUpdatedGraphics,this._addedGraphics);},performRedo:function(){var i;for(i=0;i<this._postUpdatedGraphics.length;i++){this._postUpdatedGraphics[i].setGeometry(_5.fromJson(this._postUpdatedGraphicsGeometries[i]));this._postUpdatedGraphics[i].setAttributes(this._postUpdatedGraphicsAttributes[i]);}this._featureLayer.applyEdits(this._addedGraphics,this._postUpdatedGraphics,null);}});if(_3("extend-esri")){_2.setObject("dijit.editing.Cut",_7,_4);}return _7;});