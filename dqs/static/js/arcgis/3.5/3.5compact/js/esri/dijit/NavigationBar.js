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
define("esri/dijit/NavigationBar",["dojo/_base/declare","dojo/_base/lang","dojo/_base/connect","dojo/_base/array","dojo/_base/kernel","dojo/has","dojo/query","dojo/dom","dojo/dom-class","dojo/dom-construct","esri/dijit/_TouchBase","esri/kernel"],function(_1,_2,_3,_4,_5,_6,_7,_8,_9,_a,_b,_c){var NB=_1(null,{declaredClass:"esri.dijit.NavigationBar",_items:[],constructor:function(_d,_e){var i;this.container=_8.byId(_e);this._touchBase=_b(this.container,null);this._slideDiv=_a.create("div",{},this.container,"first");this.events=[_3.connect(this._touchBase,"onclick",this,this._onClickHandler)];this._items=_d.items;_9.add(this.container,"esriMobileNavigationBar");var _f=_a.create("div",{},this._slideDiv);for(i=0;i<this._items.length;i++){var _10,div;switch(this._items[i].type){case "img":div=_a.create("div",{"class":"esriMobileNavigationItem"},_f);_10=_a.create("img",{src:this._items[i].src.toString(),style:{width:"100%",height:"100%"}},div);break;case "span":div=_a.create("div",{"class":"esriMobileNavigationItem"},_f);_10=_a.create("span",{innerHTML:this._items[i].text},div);break;case "div":div=_a.create("div",{"class":"esriMobileNavigationInfoPanel"},_f);_10=_a.create("div",{innerHTML:this._items[i].text},div);break;}_9.add(div,this._items[i].position);if(this._items[i].className){_9.add(_10,this._items[i].className);}_10._index=i;_10._item=this._items[i];this._items[i]._node=_10;}},startup:function(){this.onCreate(this._items);},destroy:function(){_4.forEach(this.events,_3.disconnect);this._touchBase=null;_5.query("img",this.container).forEach(function(_11){_11._index=null;_11._item=null;_a.destroy(_11);_11=null;});this._items=null;_a.destroy(this._slideDiv);_a.destroy(this.container);this.container=this._slideDiv=null;},getItems:function(){return this._items;},select:function(_12){this._markSelected(_12._node,_12);},onSelect:function(_13){},onUnSelect:function(_14){},onCreate:function(_15){},_onClickHandler:function(e){if(e.target.tagName.toLowerCase()==="img"){var img=e.target;var _16=img._index;var _17=img._item;_5.query("img",this.container).forEach(function(_18){if(_18!==img&&_18._item.toggleGroup===_17.toggleGroup){this._markUnSelected(_18,_18._item);}},this);this._toggleNode(img,_17);}},_toggleNode:function(_19,_1a){if(_1a.toggleState==="ON"){_1a.toggleState="OFF";if(_1a.src){_19.src=_1a.src.toString();}this.onUnSelect(_1a);}else{_1a.toggleState="ON";if(_1a.srcAlt){_19.src=_1a.srcAlt;}this.onSelect(_1a);}},_markSelected:function(_1b,_1c){_1c.toggleState="ON";if(_1c.srcAlt){_1b.src=_1c.srcAlt;}this.onSelect(_1c);},_markUnSelected:function(_1d,_1e){if(_1e.toggleState==="ON"){_1e.toggleState="OFF";if(_1e.src){_1d.src=_1e.src.toString();}this.onUnSelect(_1e);}}});if(_6("extend-esri")){_2.setObject("dijit.NavigationBar",NB,_c);}return NB;});