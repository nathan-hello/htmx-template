Goals: 
1. Groundwork: This is a template for making a web application using go, templ, htmx, and a postgres database. 
    - Notably, I don't have an opinion on what the best golang HTTP router is. This stack will use the built-in router unless it becomes too much to bear. It's up to the user to determine what HTTP router they want.
2. Make the DX akin to developing on NextJS, Astro, or any of the other great JS frameworks. 
    - Attached to this project should be a neovim lua config and VSCode settings.json for LSPs inside of .templ files that includes tailwindcss, htmx, and postgresql.
    - This also includes instructions for beginners on how to use each piece of the stack.
    - Project structure in golang, from what I hear, is a hot topic. This folder structure makes sense to me, but if it's not idiomatic to go professionals then I'm all ears to have my mind changed.
3. Minimal abstractions: Or, at least abstractions where it makes sense. Server performance is important. Abstractions should compliment (or develop!) someone's skills in standard technologies (sql, css, http, etc), not get in the way of them.
    - sqlc is a great choice for this. It makes you write actual, real-deal SQL but gives a great DX when you're inside your .go files.
4. Script the things that are annoying. This should include installing all of the relevant packages, and setting up the stack.
5. Examples: So far this includes a todo app and live chat app using websockets.
    - Auth: We're going to use Supabase Auth because it's open source, self-hostable, and an out of the box solution. It also has bucket storage, security, and a great web interface. 
        - Q: Should there be a DB layer and Auth layer? Or just DB layer?
    - Testing: A cool thing about HTMX and sqlc is that they are highly testable. Testing in web applications might be overkill, but it should be included in any examples. 
    - Cool things that are a discussion for later: SEO, SSG & SSR, Serverless, Docker containers for development and production, an optional build step for optimizing for production (re: file sizes).
        - templ has stuff in their docs regarding this. Making this work and deployable (in self-hosted, monolithic cloud, and serverless) should be a priority.

Known limitations:
1. Javascript is required: Because HTMX is client-side javascript, any browsers not using javascript can't use htmx. If you're someone who care a lot about making a JS-free webapp, then you could remove the HTMX dependency. Support for this approach is not something this project is going to consider (at least for now). Though if you're someone who cares about this, you probably already know what you're doing.
2. At time of writing, templ doesn't seem to support compilation of .templ files into a directory other than the one they are already written in. And also they will be 1:1 where `foo.templ` will go into `foo_templ.go`, opposed to having one file that takes in all of the templ. I'm sure features like these will get added over time (and I'd like to see if I could help!) but for now, the `components/` folder is a bit messy. 


The following is in order of most to least important to this stack.
1. htmx        - `https://htmx.org/docs/`
2. Postgresql  - `https://www.postgresql.org/download/`
3. sqlc        - `https://github.com/sqlc-dev/sqlc/cmd/sqlc`
4. tmpl        - `https://github.com/a-h/templ`
5. air         - `https://github.com/cosmtrek/air`
6. Tailwindcss - `https://tailwindcss.com/blog/standalone-cli`


Because this is my workflow, I'm going to assume the following about the database.
    - This project is using Supabase. - `https://supabase.com/docs/guides/cli/getting-started`
    - This project does not care about a local database (though that is a feature of supabase, and is really cool)
    - By extension, the database is Postgresql.

Environment Variables:
`$DB_URI`: This is a combination of the user, password, host, port and database name. You can find it with the Supabase web interface at Settings > Datbase > Connection String > URI.
- We need this to get the schema of the remote database with `pg_dump $DB_URI -s -f src/sqlc/input/remote-schema.sqlc`


Misc Notes:
- Turn off browser caching on your browser for your `localhost:3000` or whatever port you use. This ensures when you do a `:wa`, you'll see the change.
- To run tests, use `go test -v ./...`. This is because the tests are in the tests/ folder. Again, folder structure here is what makes sense to me, I'm open to having my mind changed. 

Problems to be ironed out:
- templ [claims to support hot reloading in the browser](https://templ.guide/commands-and-tools/hot-reload). Whenever I try to run `templ generate --proxy="localhost:3000"`, the program crashes. Maybe there's something else I should do. Even so, it's not entirely clear if the hot reloading extends to .go files. For now, we'll use air and refresh like cavemen.
  - Even using air, there is a compile time with templ or air (probably templ), where when you press refresh, the server is down until templ can compile, and air can compile that. Is there a browser setting to adjust the timeout period? Is there a misconfig somewhere? Maybe this is just the cost of using a compiled langauge, because air does compile to `tmp/` before throwing back up the server. Even so, why does it terminate? Does it know that it's not caught up yet?

VSCode
- For VSCode users, install the relevant extensions. Making a `.vscode/` folder isn't necessary because the setup is pretty easy. If this turns into a fully fledged out thing, there will be an option to add this automatically. For now, this is what you need to add. You can make them recommended extensions, then search `@recommended` in the Extensions tab to see them all, or just search for them, they're not hard to find. TODO: Provide proper workflow for working with Postgresql. 

extensions.json
```json
{ "recommendations": [
    "golang.go",
    "a-h.templ",
    "tamasfe.even-better-toml",
    "bradlc.vscode-tailwindcss",
    "mads-hartmann.bash-ide-vscode",
]}
```

settings.json
```json
{ "gopls": { "ui.semanticTokens": true } }
```
