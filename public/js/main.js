import App from './app.js';
import routes from './routes.js';
import messages from './messages.js';
import { check_login } from './service/login.js';
import { check_features, has_feature } from './service/features.js';
import { fetch_contacts, fetch_mails } from './service/mail.js';
import router_guards from './util/router_guards.js';
import { connect } from './ws.js';

function start(){
	// create router instance
	const router = VueRouter.createRouter({
		history: VueRouter.createWebHashHistory(),
		routes: routes
	});

	// set up router guards
	router_guards(router);

	const i18n = VueI18n.createI18n({
		fallbackLocale: 'en',
		messages: messages
	});

	// set up websocket events
	connect();

	// start vue
	const app = Vue.createApp(App);
	app.use(router);
	app.use(i18n);
	app.mount("#app");

	// check mails if available
	if (has_feature("mail")) {
		fetch_mails();
		fetch_contacts();
	}
}

Promise.all([
	check_login(),
	check_features()
])
.then(start);
