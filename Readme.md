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

- Uses [Go Simple Mail](https://github.com/xhit/go-simple-mail) as a 3rd-Party library for sending and receiving mails
- Uses [MailHog](https://github.com/mailhog/MailHog) as a 3rd-Party library for reading the mails.
- Uses [Foundation for Emails](https://github.com/foundation/foundation-emails) to make Email more formatted and responsive. Foundation for Emails (previously known as Ink) is a framework for creating responsive HTML emails that work in any email client — even Outlook
