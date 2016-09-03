# novo

What is Novo?

Novo is a database change management program. The purpose of the Novo program is to take a database schema file and create the necessary SQL statements to create a database. It will also be able to determine what database objects are missing from a target database, and generate the necessary SQL to create any missing objects in the target database. There will be a force and non-destructive mode. The non-destructive mode will only create objects if they do not already exist. The force mode will attempt to generate SQL that will not damage existing data, yet it will actually attempt to non-destructively morph existing matching objects to match the schema definition.

## Command List

    novo dump [host] [database] [json|sql] > [filename]

Generates a schema from an existing database. The dump format can be in JSON or SQL format.
