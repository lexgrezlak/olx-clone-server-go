BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE "message" ("id" uuid NOT NULL DEFAULT uuid_generate_v4(), "content" character varying NOT NULL, "sentTime" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP(6), "chatId" uuid, "userId" uuid, CONSTRAINT "PK_ba01f0a3e0123651915008bc578" PRIMARY KEY ("id"));
CREATE TABLE "chat" ("id" uuid NOT NULL DEFAULT uuid_generate_v4(), "postingId" uuid, CONSTRAINT "PK_9d0b2ba74336710fd31154738a5" PRIMARY KEY ("id"));
SELECT "n"."nspname", "t"."typname" FROM "pg_type" "t" INNER JOIN "pg_namespace" "n" ON "n"."oid" = "t"."typnamespace" WHERE "n"."nspname" = current_schema() AND "t"."typname" = 'posting_condition_enum';
CREATE TYPE "posting_condition_enum" AS ENUM('New', 'Used');
CREATE TABLE "posting" ("id" uuid NOT NULL DEFAULT uuid_generate_v4(), "title" character varying NOT NULL, "price" integer NOT NULL, "condition" "posting_condition_enum" NOT NULL, "description" character varying NOT NULL, "phone" integer NOT NULL, "city" character varying NOT NULL, "photos" text array NOT NULL, "createdAt" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, "updatedAt" TIMESTAMP NOT NULL DEFAULT now(), "userId" uuid, CONSTRAINT "PK_b535363e80b08416dacf34303a9" PRIMARY KEY ("id"));
CREATE TABLE "user" ("id" uuid NOT NULL DEFAULT uuid_generate_v4(), "firstName" character varying NOT NULL, "lastName" character varying NOT NULL, "email" character varying NOT NULL, "passwordHash" character varying NOT NULL, CONSTRAINT "UQ_e12875dfb3b1d92d7d7c5377e22" UNIQUE ("email"), CONSTRAINT "PK_cace4a159ff9f2512dd42373760" PRIMARY KEY ("id"));
CREATE TABLE "chat_users_user" ("chatId" uuid NOT NULL, "userId" uuid NOT NULL, CONSTRAINT "PK_c6af481280fb886733ddbd73661" PRIMARY KEY ("chatId", "userId"));
CREATE INDEX "IDX_6a573fa22dfa3574496311588c" ON "chat_users_user" ("chatId");
CREATE INDEX "IDX_2004be39e2b3044c392bfe3e61" ON "chat_users_user" ("userId");
CREATE TABLE "posting_followers_user" ("postingId" uuid NOT NULL, "userId" uuid NOT NULL, CONSTRAINT "PK_681c3d42921fd3ed36f82481369" PRIMARY KEY ("postingId", "userId"));
CREATE INDEX "IDX_a52d4af6d0a515abb7e1ca607c" ON "posting_followers_user" ("postingId");
CREATE INDEX "IDX_22a9e218f1524de474610d5efc" ON "posting_followers_user" ("userId");
ALTER TABLE "message" ADD CONSTRAINT "FK_619bc7b78eba833d2044153bacc" FOREIGN KEY ("chatId") REFERENCES "chat"("id") ON DELETE CASCADE ON UPDATE NO ACTION;
ALTER TABLE "message" ADD CONSTRAINT "FK_446251f8ceb2132af01b68eb593" FOREIGN KEY ("userId") REFERENCES "user"("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "chat" ADD CONSTRAINT "FK_c80e07229a8983e632cb84218ef" FOREIGN KEY ("postingId") REFERENCES "posting"("id") ON DELETE CASCADE ON UPDATE NO ACTION;
ALTER TABLE "posting" ADD CONSTRAINT "FK_c68c0ca3ae6b61581f640509594" FOREIGN KEY ("userId") REFERENCES "user"("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "chat_users_user" ADD CONSTRAINT "FK_6a573fa22dfa3574496311588c7" FOREIGN KEY ("chatId") REFERENCES "chat"("id") ON DELETE CASCADE ON UPDATE NO ACTION;
ALTER TABLE "chat_users_user" ADD CONSTRAINT "FK_2004be39e2b3044c392bfe3e617" FOREIGN KEY ("userId") REFERENCES "user"("id") ON DELETE CASCADE ON UPDATE NO ACTION;
ALTER TABLE "posting_followers_user" ADD CONSTRAINT "FK_a52d4af6d0a515abb7e1ca607ca" FOREIGN KEY ("postingId") REFERENCES "posting"("id") ON DELETE CASCADE ON UPDATE NO ACTION;
ALTER TABLE "posting_followers_user" ADD CONSTRAINT "FK_22a9e218f1524de474610d5efcd" FOREIGN KEY ("userId") REFERENCES "user"("id") ON DELETE CASCADE ON UPDATE NO ACTION;

END;
