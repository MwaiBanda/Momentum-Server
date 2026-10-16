-- +goose Up
--
--



SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: Meal; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public."Meal" (
    id text NOT NULL,
    reason text NOT NULL,
    email text NOT NULL,
    phone text NOT NULL,
    city text NOT NULL,
    "numOfAdults" integer NOT NULL,
    "numberOfKids" integer NOT NULL,
    street text NOT NULL,
    "preferredTime" text NOT NULL,
    favourites text NOT NULL,
    "leastFavourites" text NOT NULL,
    allergies text NOT NULL,
    instructions text NOT NULL,
    "userId" text,
    recipient text NOT NULL,
    "createdOn" timestamp(3) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


--
-- Name: MealParticipant; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public."MealParticipant" (
    "userId" text NOT NULL,
    "mealId" text NOT NULL,
    "assignedAt" timestamp(3) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    id text NOT NULL
);


--
-- Name: MealUpdates; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public."MealUpdates" (
    id text NOT NULL,
    message text NOT NULL,
    "mealId" text NOT NULL,
    "createdOn" timestamp(3) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


--
-- Name: Message; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public."Message" (
    id text NOT NULL,
    thumbnail text NOT NULL,
    series text NOT NULL,
    title text NOT NULL,
    preacher text NOT NULL,
    date text NOT NULL,
    "createdOn" timestamp(3) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    published boolean DEFAULT false NOT NULL,
    "hasOrder" boolean DEFAULT false NOT NULL
);


--
-- Name: Note; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public."Note" (
    id text NOT NULL,
    content text NOT NULL,
    "createdOn" timestamp(3) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updatedOn" timestamp(3) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "userId" text,
    "passageId" text
);


--
-- Name: Passage; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public."Passage" (
    id text NOT NULL,
    header text,
    verse text,
    message text,
    "messageId" text,
    "createdOn" timestamp(3) without time zone DEFAULT CURRENT_TIMESTAMP,
    "order" integer DEFAULT 0
);


--
-- Name: Service; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public."Service" (
    id text NOT NULL,
    name text NOT NULL,
    type text NOT NULL,
    "Image" text NOT NULL,
    "createdOn" timestamp(3) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


--
-- Name: Transaction; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public."Transaction" (
    id text NOT NULL,
    amount integer NOT NULL,
    date text NOT NULL,
    "createdOn" timestamp(3) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "userId" text NOT NULL,
    "Description" text NOT NULL
);


--
-- Name: User; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public."User" (
    id text NOT NULL,
    email text NOT NULL,
    fullname text NOT NULL,
    phone text NOT NULL,
    "createdOn" timestamp(3) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


--
-- Name: VolunteerService; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public."VolunteerService" (
    id text NOT NULL,
    description text NOT NULL,
    email text NOT NULL,
    phone text NOT NULL,
    organizer text NOT NULL,
    "serviceId" text,
    "userId" text,
    "createdOn" timestamp(3) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


--
-- Name: VolunteeredMeal; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public."VolunteeredMeal" (
    id text NOT NULL,
    description text NOT NULL,
    notes text NOT NULL,
    "userId" text,
    "mealId" text,
    "createdOn" timestamp(3) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    date text DEFAULT ''::text NOT NULL
);


--
-- Name: VolunteeredServiceDay; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public."VolunteeredServiceDay" (
    id text NOT NULL,
    notes text NOT NULL,
    "userId" text,
    "serviceId" text NOT NULL,
    "createdOn" timestamp(3) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    date text DEFAULT ''::text NOT NULL
);


--
-- Name: VolunteeredServiceParticipant; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public."VolunteeredServiceParticipant" (
    "userId" text NOT NULL,
    "serviceId" text NOT NULL,
    "assignedAt" timestamp(3) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    id text NOT NULL
);


--
-- Name: MealParticipant MealParticipant_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public."MealParticipant"
    ADD CONSTRAINT "MealParticipant_pkey" PRIMARY KEY (id, "userId", "mealId");


--
-- Name: MealUpdates MealUpdates_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public."MealUpdates"
    ADD CONSTRAINT "MealUpdates_pkey" PRIMARY KEY (id);


--
-- Name: Meal Meal_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public."Meal"
    ADD CONSTRAINT "Meal_pkey" PRIMARY KEY (id);


--
-- Name: Message Message_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public."Message"
    ADD CONSTRAINT "Message_pkey" PRIMARY KEY (id);


--
-- Name: Note Note_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public."Note"
    ADD CONSTRAINT "Note_pkey" PRIMARY KEY (id);


--
-- Name: Passage Passage_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public."Passage"
    ADD CONSTRAINT "Passage_pkey" PRIMARY KEY (id);


--
-- Name: Service Service_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public."Service"
    ADD CONSTRAINT "Service_pkey" PRIMARY KEY (id);


--
-- Name: Transaction Transaction_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public."Transaction"
    ADD CONSTRAINT "Transaction_pkey" PRIMARY KEY (id);


--
-- Name: User User_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public."User"
    ADD CONSTRAINT "User_pkey" PRIMARY KEY (id);


--
-- Name: VolunteerService VolunteerService_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public."VolunteerService"
    ADD CONSTRAINT "VolunteerService_pkey" PRIMARY KEY (id);


--
-- Name: VolunteeredMeal VolunteeredMeal_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public."VolunteeredMeal"
    ADD CONSTRAINT "VolunteeredMeal_pkey" PRIMARY KEY (id);


--
-- Name: VolunteeredServiceDay VolunteeredServiceDay_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public."VolunteeredServiceDay"
    ADD CONSTRAINT "VolunteeredServiceDay_pkey" PRIMARY KEY (id);


--
-- Name: VolunteeredServiceParticipant VolunteeredServiceParticipant_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public."VolunteeredServiceParticipant"
    ADD CONSTRAINT "VolunteeredServiceParticipant_pkey" PRIMARY KEY (id, "userId", "serviceId");


--
-- Name: MealUpdates_mealId_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX "MealUpdates_mealId_idx" ON public."MealUpdates" USING btree ("mealId");


--
-- Name: Meal_userId_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX "Meal_userId_idx" ON public."Meal" USING btree ("userId");


--
-- Name: Note_userId_passageId_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX "Note_userId_passageId_idx" ON public."Note" USING btree ("userId", "passageId");


--
-- Name: Passage_messageId_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX "Passage_messageId_idx" ON public."Passage" USING btree ("messageId");


--
-- Name: Service_type_key; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX "Service_type_key" ON public."Service" USING btree (type);


--
-- Name: Transaction_userId_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX "Transaction_userId_idx" ON public."Transaction" USING btree ("userId");


--
-- Name: VolunteerService_userId_serviceId_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX "VolunteerService_userId_serviceId_idx" ON public."VolunteerService" USING btree ("userId", "serviceId");


--
-- Name: VolunteeredMeal_userId_mealId_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX "VolunteeredMeal_userId_mealId_idx" ON public."VolunteeredMeal" USING btree ("userId", "mealId");


--
-- Name: VolunteeredServiceDay_userId_serviceId_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX "VolunteeredServiceDay_userId_serviceId_idx" ON public."VolunteeredServiceDay" USING btree ("userId", "serviceId");


--
-- Name: MealParticipant MealParticipant_mealId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public."MealParticipant"
    ADD CONSTRAINT "MealParticipant_mealId_fkey" FOREIGN KEY ("mealId") REFERENCES public."Meal"(id) ON UPDATE CASCADE ON DELETE RESTRICT;


--
-- Name: MealParticipant MealParticipant_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public."MealParticipant"
    ADD CONSTRAINT "MealParticipant_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."User"(id) ON UPDATE CASCADE ON DELETE RESTRICT;


--
-- Name: MealUpdates MealUpdates_mealId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public."MealUpdates"
    ADD CONSTRAINT "MealUpdates_mealId_fkey" FOREIGN KEY ("mealId") REFERENCES public."VolunteeredMeal"(id) ON UPDATE CASCADE ON DELETE RESTRICT;


--
-- Name: Meal Meal_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public."Meal"
    ADD CONSTRAINT "Meal_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."User"(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- Name: Note Note_passageId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public."Note"
    ADD CONSTRAINT "Note_passageId_fkey" FOREIGN KEY ("passageId") REFERENCES public."Passage"(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- Name: Note Note_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public."Note"
    ADD CONSTRAINT "Note_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."User"(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- Name: Passage Passage_messageId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public."Passage"
    ADD CONSTRAINT "Passage_messageId_fkey" FOREIGN KEY ("messageId") REFERENCES public."Message"(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- Name: Transaction Transaction_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public."Transaction"
    ADD CONSTRAINT "Transaction_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."User"(id) ON UPDATE CASCADE ON DELETE RESTRICT;


--
-- Name: VolunteerService VolunteerService_serviceId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public."VolunteerService"
    ADD CONSTRAINT "VolunteerService_serviceId_fkey" FOREIGN KEY ("serviceId") REFERENCES public."Service"(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- Name: VolunteerService VolunteerService_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public."VolunteerService"
    ADD CONSTRAINT "VolunteerService_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."User"(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- Name: VolunteeredMeal VolunteeredMeal_mealId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public."VolunteeredMeal"
    ADD CONSTRAINT "VolunteeredMeal_mealId_fkey" FOREIGN KEY ("mealId") REFERENCES public."Meal"(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- Name: VolunteeredMeal VolunteeredMeal_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public."VolunteeredMeal"
    ADD CONSTRAINT "VolunteeredMeal_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."User"(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- Name: VolunteeredServiceDay VolunteeredServiceDay_serviceId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public."VolunteeredServiceDay"
    ADD CONSTRAINT "VolunteeredServiceDay_serviceId_fkey" FOREIGN KEY ("serviceId") REFERENCES public."VolunteerService"(id) ON UPDATE CASCADE ON DELETE RESTRICT;


--
-- Name: VolunteeredServiceDay VolunteeredServiceDay_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public."VolunteeredServiceDay"
    ADD CONSTRAINT "VolunteeredServiceDay_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."User"(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- Name: VolunteeredServiceParticipant VolunteeredServiceParticipant_serviceId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public."VolunteeredServiceParticipant"
    ADD CONSTRAINT "VolunteeredServiceParticipant_serviceId_fkey" FOREIGN KEY ("serviceId") REFERENCES public."VolunteerService"(id) ON UPDATE CASCADE ON DELETE RESTRICT;


--
-- Name: VolunteeredServiceParticipant VolunteeredServiceParticipant_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public."VolunteeredServiceParticipant"
    ADD CONSTRAINT "VolunteeredServiceParticipant_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."User"(id) ON UPDATE CASCADE ON DELETE RESTRICT;


--
--

SELECT pg_catalog.set_config('search_path', 'public', false);

-- +goose Down
DROP TABLE IF EXISTS public."VolunteeredServiceParticipant";
DROP TABLE IF EXISTS public."VolunteeredServiceDay";
DROP TABLE IF EXISTS public."VolunteeredMeal";
DROP TABLE IF EXISTS public."MealUpdates";
DROP TABLE IF EXISTS public."MealParticipant";
DROP TABLE IF EXISTS public."Note";
DROP TABLE IF EXISTS public."Passage";
DROP TABLE IF EXISTS public."VolunteerService";
DROP TABLE IF EXISTS public."Transaction";
DROP TABLE IF EXISTS public."Meal";
DROP TABLE IF EXISTS public."Message";
DROP TABLE IF EXISTS public."Service";
DROP TABLE IF EXISTS public."User";
