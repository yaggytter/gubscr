
function find2(s) {


$.get("/gitrepos/yaggytter/embulk-input-slack-history/master/embulk-input-slack-history/lib/embulk/input/slack_history.rb", function(data){
  $('#code1').html(data);

  x=new RegExp(s,'gim');
  rn=Math.floor(Math.random()*100);
  rid='z' + rn;
  b = document.body.innerHTML;

  console.log(b);

  b.match(/^func/igm);
  matchstr=RegExp.lastMatch;
  console.log(matchstr);
  b=b.replace(x,'<span name=' + rid + ' id=' + rid + ' style=\'color:#000; background-color:yellow; font-weight:bold;\'>'+matchstr+'</span>');

  void(document.body.innerHTML=b);

//  alert('Found ' + document.getElementsByName(rid).length + ' matches.');

  window.scrollTo(0,document.getElementsByName(rid)[0].offsetTop);

$('pre code').each(function(i, block) {
    hljs.highlightBlock(block);
  });

});


}


function getSelectedText() {
      if('\v' === 'v')
        return document.selection.createRange().text;
      else
        return ('getSelection' in window ? window : document).getSelection().toString();
 }

$(function() {

	var interval = 720;

	shortcut.add("Ctrl+Shift+F", function() {
		find2('^func');
//		alert(getSelectedText());
	});

	shortcut.add("Ctrl+Shift+B", function() {
		alert("Hi there!");
	});

	$("#push").bind("touchend", function() {
		find("func");
		alert(getSelectedText());
    });

});
