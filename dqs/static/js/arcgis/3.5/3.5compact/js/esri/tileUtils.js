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
define("esri/tileUtils",["dojo/_base/array","dojo/has","esri/kernel","esri/geometry/Point","esri/geometry/Extent"],function(_1,_2,_3,_4,_5){function _6(_7,ti,_8){var wd=_7.width,ht=_7.height,ew=_8.xmax-_8.xmin,eh=_8.ymax-_8.ymin,_9=(_7.__tileInfo===ti),_a=_9?_7.getMinZoom():-1,_b=_9?_7.getMaxZoom():-1,ed=-1,_c=ti.lods,i,_d=Math.abs,_e,cl,_f;_a=(_a>-1)?_a:0;_b=(_b>-1)?_b:(_c.length-1);for(i=_a;i<=_b;i++){cl=_c[i];if(!cl){continue;}_f=ew>eh?_d(eh-(ht*cl.resolution)):_d(ew-(wd*cl.resolution));if(ed<0||_f<=ed){_e=cl;ed=_f;}else{break;}}return _e;};function _10(map,_11,lod){var res=lod.resolution,cx=(_11.xmin+_11.xmax)/2,cy=(_11.ymin+_11.ymax)/2,_12=(map.width/2)*res,_13=(map.height/2)*res;return new _5(cx-_12,cy-_13,cx+_12,cy+_13,_11.spatialReference);};function _14(map,ti,_15,lod){var res=lod.resolution,tw=ti.width,th=ti.height,to=ti.origin,mv=map.__visibleDelta,_16=Math.floor,tmw=tw*res,tmh=th*res,tr=_16((to.y-_15.y)/tmh),tc=_16((_15.x-to.x)/tmw),_17=to.x+(tc*tmw),_18=to.y-(tr*tmh),oX=_16(Math.abs((_15.x-_17)*tw/tmw))+mv.x,oY=_16(Math.abs((_15.y-_18)*th/tmh))+mv.y;return {point:_15,coords:{row:tr,col:tc},offsets:{x:oX,y:oY}};};var _19={_addFrameInfo:function(_1a,_1b){var _1c,_1d,_1e=2*_1b.origin[1],_1f=_1b.origin[0],_20=_1a.origin.x,_21=_1a.width,_22;_1.forEach(_1a.lods,function(lod){_1c=Math.round(_1e/lod.resolution);_1d=Math.ceil(_1c/_21);_22=Math.floor((_1f-_20)/(_21*lod.resolution));if(!lod._frameInfo){lod._frameInfo=[_1d,_22,_22+_1d-1,_1c];}});},getContainingTileCoords:function(ti,_23,lod){var to=ti.origin,res=lod.resolution,tmw=ti.width*res,tmh=ti.height*res,tc=Math.floor((_23.x-to.x)/tmw),tr=Math.floor((to.y-_23.y)/tmh);return {row:tr,col:tc};},getCandidateTileInfo:function(map,ti,_24){var lod=_6(map,ti,_24),adj=_10(map,_24,lod),ct=_14(map,ti,new _4(adj.xmin,adj.ymax,_24.spatialReference),lod);return {tile:ct,lod:lod,extent:adj};},getTileExtent:function(ti,_25,row,col){var to=ti.origin,lod=ti.lods[_25],res=lod.resolution,tw=ti.width,th=ti.height;return new _5(((col*res)*tw)+to.x,to.y-((row+1)*res)*th,(((col+1)*res)*tw)+to.x,to.y-((row*res)*th),ti.spatialReference);}};if(_2("extend-esri")){_3.TileUtils=_19;}return _19;});