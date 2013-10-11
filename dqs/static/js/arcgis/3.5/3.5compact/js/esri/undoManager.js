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
define("esri/undoManager",["dojo/_base/declare","dojo/has","esri/kernel","dojo/has!extend-esri?esri/OperationBase"],function(_1,_2,_3){var _4=_1(null,{declaredClass:"esri.UndoManager",maxOperations:10,canUndo:false,canRedo:false,position:0,length:0,onUndo:function(){},onRedo:function(){},onAdd:function(){},onChange:function(){},constructor:function(_5){_5=_5||{};if(_5.maxOperations){this.maxOperations=_5.maxOperations;}this._historyStack=[];},add:function(_6){if(this.maxOperations>0){while(this._historyStack.length>=this.maxOperations){this._historyStack.shift();}}this._historyStack.splice(this.position,0,_6);this.position++;this.clearRedo();this.onAdd();this._checkAvailability();},undo:function(){if(this.position===0){return null;}var _7=this.peekUndo();this.position--;if(_7){_7.performUndo();}this.onUndo();this._checkAvailability();},redo:function(){if(this.position===this._historyStack.length){return null;}var _8=this.peekRedo();this.position++;if(_8){_8.performRedo();}this.onRedo();this._checkAvailability();},_checkAvailability:function(){this.length=this._historyStack.length;if(this.length===0){this.canRedo=false;this.canUndo=false;}else{if(this.position===0){this.canRedo=true;this.canUndo=false;}else{if(this.position===this.length){this.canUndo=true;this.canRedo=false;}else{this.canUndo=true;this.canRedo=true;}}}this.onChange();},clearUndo:function(){this._historyStack.splice(0,this.position);this.position=0;this._checkAvailability();},clearRedo:function(){this._historyStack.splice(this.position,this._historyStack.length-this.position);this.position=this._historyStack.length;this._checkAvailability();},peekUndo:function(){if(this._historyStack.length>0&&this.position>0){return this.get(this.position-1);}},peekRedo:function(){if(this._historyStack.length>0&&this.position<this._historyStack.length){return this.get(this.position);}},get:function(_9){return this._historyStack[_9];},remove:function(_a){if(this._historyStack.length>0){this._historyStack.splice(_a,1);if(this.position>0){if(_a<this.position){this.position--;}}this._checkAvailability();}},destroy:function(){this._historyStack=null;}});if(_2("extend-esri")){_3.UndoManager=_4;}return _4;});