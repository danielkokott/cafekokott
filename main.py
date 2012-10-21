import webapp2
import jinja2
import os

jinja_environment = jinja2.Environment(
    loader=jinja2.FileSystemLoader(os.path.dirname(__file__)))

class MainPage(webapp2.RequestHandler):
    def get(self):
        # guestbook_name=self.request.get('guestbook_name')
        # greetings_query = Greeting.all().ancestor(
        #     guestbook_key(guestbook_name)).order('-date')
        # greetings = greetings_query.fetch(10)

        # if users.get_current_user():
        #     url = users.create_logout_url(self.request.uri)
        #     url_linktext = 'Logout'
        # else:
        #     url = users.create_login_url(self.request.uri)
        #     url_linktext = 'Login'

        # template_values = {
        #     'greetings': greetings,
        #     'url': url,
        #     'url_linktext': url_linktext,
        # }

        template = jinja_environment.get_template('web/workinprogress.html')
        self.response.out.write(template.render())

app = webapp2.WSGIApplication([('/', MainPage)],
                              debug=True)