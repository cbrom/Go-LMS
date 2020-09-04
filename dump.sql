--
-- PostgreSQL database dump
--

-- Dumped from database version 12.3
-- Dumped by pg_dump version 12.3

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: answer_options; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.answer_options (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    quiz_question_id uuid,
    value text,
    hint character varying(255)
);


ALTER TABLE public.answer_options OWNER TO postgres;

--
-- Name: certificates; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.certificates (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    course_id uuid,
    course_author_id uuid,
    qr_corner text,
    qr_scale integer,
    margin integer,
    name_offset_top integer,
    font_size integer,
    message text,
    active boolean
);


ALTER TABLE public.certificates OWNER TO postgres;

--
-- Name: content_blocks; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.content_blocks (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    block_type character varying(100),
    content jsonb,
    sort_index integer,
    target_version_id uuid
);


ALTER TABLE public.content_blocks OWNER TO postgres;

--
-- Name: course_authors; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.course_authors (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    user_id uuid,
    course_id uuid
);


ALTER TABLE public.course_authors OWNER TO postgres;

--
-- Name: courses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.courses (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name character varying(100),
    ends_at timestamp with time zone,
    description text,
    enable_leadboard boolean,
    public_signup boolean,
    featured boolean,
    about character varying(100),
    progression_behavior character varying(100),
    progression_limit integer
);


ALTER TABLE public.courses OWNER TO postgres;

--
-- Name: evaluation_criterias; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.evaluation_criterias (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name character varying(100),
    course_id uuid,
    max_grade integer,
    pass_grade integer,
    grade_labels jsonb
);


ALTER TABLE public.evaluation_criterias OWNER TO postgres;

--
-- Name: issued_certificates; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.issued_certificates (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    certificate_id uuid,
    user_id uuid,
    serial_number text
);


ALTER TABLE public.issued_certificates OWNER TO postgres;

--
-- Name: levels; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.levels (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name character varying(100),
    course_id uuid,
    description text,
    number integer,
    unlock_on timestamp with time zone
);


ALTER TABLE public.levels OWNER TO postgres;

--
-- Name: quiz_question_user_answers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.quiz_question_user_answers (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    question_id uuid,
    answer_id uuid,
    user_id uuid
);


ALTER TABLE public.quiz_question_user_answers OWNER TO postgres;

--
-- Name: quiz_questions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.quiz_questions (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    quiz_id uuid,
    question text,
    description character varying(100),
    correct_answer_id text
);


ALTER TABLE public.quiz_questions OWNER TO postgres;

--
-- Name: quizzes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.quizzes (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    title text,
    target_id uuid
);


ALTER TABLE public.quizzes OWNER TO postgres;

--
-- Name: student_courses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.student_courses (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    user_id uuid,
    course_id uuid
);


ALTER TABLE public.student_courses OWNER TO postgres;

--
-- Name: target_groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.target_groups (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name character varying(100),
    description text,
    sort_index integer,
    milestone boolean,
    level_id uuid
);


ALTER TABLE public.target_groups OWNER TO postgres;

--
-- Name: target_versions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.target_versions (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    target_id uuid,
    version_name character varying(100)
);


ALTER TABLE public.target_versions OWNER TO postgres;

--
-- Name: targets; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.targets (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    role character varying(100),
    title character varying(100),
    description text,
    completion_instructions text,
    resource_url character varying(255),
    target_group_id uuid,
    sort_index integer,
    session_at timestamp with time zone,
    link_to_complete character varying(255),
    resubmittable boolean,
    check_list jsonb,
    review_checklist jsonb
);


ALTER TABLE public.targets OWNER TO postgres;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    email character varying(100),
    password_salt text,
    password_hash bytea,
    role integer,
    sigin_in_count integer,
    current_sign_in_at timestamp with time zone,
    last_sign_in_at timestamp with time zone,
    current_sign_in_ip text,
    last_sign_in_ip text,
    remember_token text,
    confirmed_at timestamp with time zone,
    confirmation_mail_sent_at timestamp with time zone,
    name text,
    phone text,
    title text,
    key_skills text,
    about text,
    time_zone timestamp with time zone
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: answer_options answer_options_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.answer_options
    ADD CONSTRAINT answer_options_pkey PRIMARY KEY (id);


--
-- Name: certificates certificates_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.certificates
    ADD CONSTRAINT certificates_pkey PRIMARY KEY (id);


--
-- Name: content_blocks content_blocks_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.content_blocks
    ADD CONSTRAINT content_blocks_pkey PRIMARY KEY (id);


--
-- Name: course_authors course_authors_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.course_authors
    ADD CONSTRAINT course_authors_pkey PRIMARY KEY (id);


--
-- Name: courses courses_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.courses
    ADD CONSTRAINT courses_name_key UNIQUE (name);


--
-- Name: courses courses_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.courses
    ADD CONSTRAINT courses_pkey PRIMARY KEY (id);


--
-- Name: evaluation_criterias evaluation_criterias_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.evaluation_criterias
    ADD CONSTRAINT evaluation_criterias_pkey PRIMARY KEY (id);


--
-- Name: issued_certificates issued_certificates_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.issued_certificates
    ADD CONSTRAINT issued_certificates_pkey PRIMARY KEY (id);


--
-- Name: levels levels_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.levels
    ADD CONSTRAINT levels_pkey PRIMARY KEY (id);


--
-- Name: quiz_question_user_answers quiz_question_user_answers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.quiz_question_user_answers
    ADD CONSTRAINT quiz_question_user_answers_pkey PRIMARY KEY (id);


--
-- Name: quiz_questions quiz_questions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.quiz_questions
    ADD CONSTRAINT quiz_questions_pkey PRIMARY KEY (id);


--
-- Name: quizzes quizzes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.quizzes
    ADD CONSTRAINT quizzes_pkey PRIMARY KEY (id);


--
-- Name: student_courses student_courses_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.student_courses
    ADD CONSTRAINT student_courses_pkey PRIMARY KEY (id);


--
-- Name: target_groups target_groups_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.target_groups
    ADD CONSTRAINT target_groups_pkey PRIMARY KEY (id);


--
-- Name: target_versions target_versions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.target_versions
    ADD CONSTRAINT target_versions_pkey PRIMARY KEY (id);


--
-- Name: targets targets_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.targets
    ADD CONSTRAINT targets_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: idx_answer_options_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_answer_options_deleted_at ON public.answer_options USING btree (deleted_at);


--
-- Name: idx_certificates_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_certificates_deleted_at ON public.certificates USING btree (deleted_at);


--
-- Name: idx_content_blocks_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_content_blocks_deleted_at ON public.content_blocks USING btree (deleted_at);


--
-- Name: idx_course_authors_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_course_authors_deleted_at ON public.course_authors USING btree (deleted_at);


--
-- Name: idx_courses_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_courses_deleted_at ON public.courses USING btree (deleted_at);


--
-- Name: idx_evaluation_criterias_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_evaluation_criterias_deleted_at ON public.evaluation_criterias USING btree (deleted_at);


--
-- Name: idx_issued_certificates_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_issued_certificates_deleted_at ON public.issued_certificates USING btree (deleted_at);


--
-- Name: idx_levels_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_levels_deleted_at ON public.levels USING btree (deleted_at);


--
-- Name: idx_quiz_question_user_answers_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_quiz_question_user_answers_deleted_at ON public.quiz_question_user_answers USING btree (deleted_at);


--
-- Name: idx_quiz_questions_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_quiz_questions_deleted_at ON public.quiz_questions USING btree (deleted_at);


--
-- Name: idx_quizzes_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_quizzes_deleted_at ON public.quizzes USING btree (deleted_at);


--
-- Name: idx_student_courses_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_student_courses_deleted_at ON public.student_courses USING btree (deleted_at);


--
-- Name: idx_target_groups_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_target_groups_deleted_at ON public.target_groups USING btree (deleted_at);


--
-- Name: idx_target_versions_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_target_versions_deleted_at ON public.target_versions USING btree (deleted_at);


--
-- Name: idx_targets_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_targets_deleted_at ON public.targets USING btree (deleted_at);


--
-- Name: idx_users_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_users_deleted_at ON public.users USING btree (deleted_at);


--
-- Name: uix_users_email; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX uix_users_email ON public.users USING btree (email);


--
-- Name: answer_options answer_options_quiz_question_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.answer_options
    ADD CONSTRAINT answer_options_quiz_question_id_fkey FOREIGN KEY (quiz_question_id) REFERENCES public.quiz_questions(id) ON DELETE CASCADE;


--
-- Name: certificates certificates_course_author_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.certificates
    ADD CONSTRAINT certificates_course_author_id_fkey FOREIGN KEY (course_author_id) REFERENCES public.course_authors(id) ON DELETE CASCADE;


--
-- Name: certificates certificates_course_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.certificates
    ADD CONSTRAINT certificates_course_id_fkey FOREIGN KEY (course_id) REFERENCES public.courses(id) ON DELETE CASCADE;


--
-- Name: content_blocks content_blocks_target_version_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.content_blocks
    ADD CONSTRAINT content_blocks_target_version_id_fkey FOREIGN KEY (target_version_id) REFERENCES public.target_versions(id) ON DELETE CASCADE;


--
-- Name: course_authors course_authors_course_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.course_authors
    ADD CONSTRAINT course_authors_course_id_fkey FOREIGN KEY (course_id) REFERENCES public.courses(id) ON DELETE CASCADE;


--
-- Name: course_authors course_authors_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.course_authors
    ADD CONSTRAINT course_authors_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: evaluation_criterias evaluation_criterias_course_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.evaluation_criterias
    ADD CONSTRAINT evaluation_criterias_course_id_fkey FOREIGN KEY (course_id) REFERENCES public.courses(id) ON DELETE CASCADE;


--
-- Name: issued_certificates issued_certificates_certificate_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.issued_certificates
    ADD CONSTRAINT issued_certificates_certificate_id_fkey FOREIGN KEY (certificate_id) REFERENCES public.certificates(id) ON DELETE CASCADE;


--
-- Name: issued_certificates issued_certificates_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.issued_certificates
    ADD CONSTRAINT issued_certificates_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: levels levels_course_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.levels
    ADD CONSTRAINT levels_course_id_fkey FOREIGN KEY (course_id) REFERENCES public.courses(id) ON DELETE CASCADE;


--
-- Name: quiz_question_user_answers quiz_question_user_answers_answer_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.quiz_question_user_answers
    ADD CONSTRAINT quiz_question_user_answers_answer_id_fkey FOREIGN KEY (answer_id) REFERENCES public.answer_options(id) ON DELETE CASCADE;


--
-- Name: quiz_question_user_answers quiz_question_user_answers_question_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.quiz_question_user_answers
    ADD CONSTRAINT quiz_question_user_answers_question_id_fkey FOREIGN KEY (question_id) REFERENCES public.quiz_questions(id) ON DELETE CASCADE;


--
-- Name: quiz_question_user_answers quiz_question_user_answers_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.quiz_question_user_answers
    ADD CONSTRAINT quiz_question_user_answers_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: quiz_questions quiz_questions_quiz_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.quiz_questions
    ADD CONSTRAINT quiz_questions_quiz_id_fkey FOREIGN KEY (quiz_id) REFERENCES public.quizzes(id) ON DELETE CASCADE;


--
-- Name: quizzes quizzes_target_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.quizzes
    ADD CONSTRAINT quizzes_target_id_fkey FOREIGN KEY (target_id) REFERENCES public.targets(id) ON DELETE CASCADE;


--
-- Name: student_courses student_courses_course_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.student_courses
    ADD CONSTRAINT student_courses_course_id_fkey FOREIGN KEY (course_id) REFERENCES public.courses(id) ON DELETE CASCADE;


--
-- Name: student_courses student_courses_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.student_courses
    ADD CONSTRAINT student_courses_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: target_groups target_groups_level_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.target_groups
    ADD CONSTRAINT target_groups_level_id_fkey FOREIGN KEY (level_id) REFERENCES public.levels(id) ON DELETE CASCADE;


--
-- Name: target_versions target_versions_target_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.target_versions
    ADD CONSTRAINT target_versions_target_id_fkey FOREIGN KEY (target_id) REFERENCES public.targets(id) ON DELETE CASCADE;


--
-- Name: targets targets_target_group_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.targets
    ADD CONSTRAINT targets_target_group_id_fkey FOREIGN KEY (target_group_id) REFERENCES public.target_groups(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

