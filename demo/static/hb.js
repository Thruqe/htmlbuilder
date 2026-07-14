// htmlbuilder runtime — auto-included when any node uses OnClick.
// Do not include this manually; Document.Render wires it in for you.

function __hb_call(url, targetId) {
	fetch(url, { method: 'POST' })
		.then(function (res) { return res.text(); })
		.then(function (html) {
			document.getElementById(targetId).outerHTML = html;
		});
}

function __hb_toggle_theme(iconSelector) {
	var html = document.documentElement;
	var current = html.getAttribute('data-theme');
	var next = current === 'dark' ? 'light' : 'dark';
	html.setAttribute('data-theme', next);

	var icon = document.querySelector(iconSelector);
	if (icon) {
		icon.className = next === 'dark' ? 'fi fi-rr-sun' : 'fi fi-rr-moon';
	}
}