# Bookings and reservations

This is the repository for the bookings and reservations project.

- Built in GO version 1.19.
- Uses the [chi router](http://github.com/go-chi/chi).
- Uses the [alex edwards SCS](http://github.com/alexedwards/scs) session management.
- Uses [nosurf](http://github.com/justinas/nosurf).

# Bookings Layout

Created the bookings and reservations layout using HTML5 and Vanilla JS.

- Uses [Bootstrap](https://getbootstrap.com/) for styling.
- Uses [DatePicker](https://github.com/mymth/vanillajs-datepicker) for selecting Arrival/Departure dates.
- Diplaying notifications using [Notie](https://github.com/jaredreich/notie).
- Uses [Sweetalert2](https://github.com/sweetalert2/sweetalert2) for creating the modals.
- Uses [govalidator](https://github.com/asaskevich/govalidator) for validating form module.

# Database

Using the Postgres database and DBeaver for interaction.

- Uses [Postgres](https://github.com/postgres/postgres) database for storing the user records and schema.
- Uses [Soda](https://github.com/gobuffalo/soda) for creating schemas and yml files and interacting with them.
- For GUI side it is using [DBeaver](https://github.com/dbeaver/dbeaver) for checking up the schemas and records in database.

# Mailings

Using some standard Built-in and 3rd-Party libraries for sending/receiving mails

- Uses [Go Simple Mail](https://github.com/xhit/go-simple-mail) as a 3rd-Party library for sending and receiving mails.
- Uses [MailHog](https://github.com/mailhog/MailHog) as a 3rd-Party library for reading the mails.
- Uses [Foundation for Emails](https://github.com/foundation/foundation-emails) to make Email more formatted and responsive. Foundation for Emails (previously known as Ink) is a framework for creating responsive HTML emails that work in any email client — even Outlook.

# Admin Dashboard

Creating a admin dashboard for handling all the administrative tasks

- Uses [RoyalUI-Free-Bootstrap-Admin-Template](https://github.com/BootstrapDash/RoyalUI-Free-Bootstrap-Admin-Template) RoyalUI is a highly responsive template built with the latest version of Bootstrap, CSS, HTML5, jQuery, and SASS it provides plenty of handy Dashboard elements, useful tools, and other features using this template.
- Uses [Simple-DataTables](https://github.com/fiduswriter/Simple-DataTables) for creating reservations tables and displaying data in tabular format it is a lightweight, extendable, dependency-free javascript HTML table plugin. Similar to jQuery DataTables for use in modern browsers, but without the jQuery dependency.

# Production

While making your site live we don't want some creditianls to be distributed locally.

- Create an .sh/.bat files for executing cmd based codes
- Use [GoDotEnv](https://github.com/joho/godotenv)which is made for Go (golang) port of the Ruby dotenv project (which loads env vars from a .env file).
