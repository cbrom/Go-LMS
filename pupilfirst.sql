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
-- Name: citext; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS citext WITH SCHEMA public;


--
-- Name: EXTENSION citext; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION citext IS 'data type for case-insensitive character strings';


--
-- Name: pg_stat_statements; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS pg_stat_statements WITH SCHEMA public;


--
-- Name: EXTENSION pg_stat_statements; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION pg_stat_statements IS 'track execution statistics of all SQL statements executed';


--
-- Name: pg_trgm; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS pg_trgm WITH SCHEMA public;


--
-- Name: EXTENSION pg_trgm; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION pg_trgm IS 'text similarity measurement and index searching based on trigrams';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: active_admin_comments; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.active_admin_comments (
    id integer NOT NULL,
    namespace character varying,
    body text,
    resource_type character varying NOT NULL,
    author_id integer,
    author_type character varying,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    resource_id integer NOT NULL
);


ALTER TABLE public.active_admin_comments OWNER TO mac;

--
-- Name: active_admin_comments_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.active_admin_comments_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.active_admin_comments_id_seq OWNER TO mac;

--
-- Name: active_admin_comments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.active_admin_comments_id_seq OWNED BY public.active_admin_comments.id;


--
-- Name: active_storage_attachments; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.active_storage_attachments (
    id bigint NOT NULL,
    name character varying NOT NULL,
    record_type character varying NOT NULL,
    record_id bigint NOT NULL,
    blob_id bigint NOT NULL,
    created_at timestamp without time zone NOT NULL
);


ALTER TABLE public.active_storage_attachments OWNER TO mac;

--
-- Name: active_storage_attachments_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.active_storage_attachments_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.active_storage_attachments_id_seq OWNER TO mac;

--
-- Name: active_storage_attachments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.active_storage_attachments_id_seq OWNED BY public.active_storage_attachments.id;


--
-- Name: active_storage_blobs; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.active_storage_blobs (
    id bigint NOT NULL,
    key character varying NOT NULL,
    filename character varying NOT NULL,
    content_type character varying,
    metadata text,
    byte_size bigint NOT NULL,
    checksum character varying NOT NULL,
    created_at timestamp without time zone NOT NULL
);


ALTER TABLE public.active_storage_blobs OWNER TO mac;

--
-- Name: active_storage_blobs_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.active_storage_blobs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.active_storage_blobs_id_seq OWNER TO mac;

--
-- Name: active_storage_blobs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.active_storage_blobs_id_seq OWNED BY public.active_storage_blobs.id;


--
-- Name: admin_users; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.admin_users (
    id integer NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    username character varying,
    fullname character varying,
    user_id integer,
    email character varying
);


ALTER TABLE public.admin_users OWNER TO mac;

--
-- Name: admin_users_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.admin_users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.admin_users_id_seq OWNER TO mac;

--
-- Name: admin_users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.admin_users_id_seq OWNED BY public.admin_users.id;


--
-- Name: ahoy_events; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.ahoy_events (
    id uuid NOT NULL,
    visit_id uuid,
    user_id integer,
    user_type character varying,
    name character varying,
    properties jsonb,
    "time" timestamp without time zone
);


ALTER TABLE public.ahoy_events OWNER TO mac;

--
-- Name: answer_options; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.answer_options (
    id bigint NOT NULL,
    quiz_question_id bigint,
    value text,
    hint text,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.answer_options OWNER TO mac;

--
-- Name: answer_options_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.answer_options_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.answer_options_id_seq OWNER TO mac;

--
-- Name: answer_options_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.answer_options_id_seq OWNED BY public.answer_options.id;


--
-- Name: applicants; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.applicants (
    id bigint NOT NULL,
    email character varying,
    name character varying,
    login_token character varying,
    login_mail_sent_at timestamp without time zone,
    course_id bigint,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.applicants OWNER TO mac;

--
-- Name: applicants_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.applicants_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.applicants_id_seq OWNER TO mac;

--
-- Name: applicants_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.applicants_id_seq OWNED BY public.applicants.id;


--
-- Name: ar_internal_metadata; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.ar_internal_metadata (
    key character varying NOT NULL,
    value character varying,
    created_at timestamp(6) without time zone NOT NULL,
    updated_at timestamp(6) without time zone NOT NULL
);


ALTER TABLE public.ar_internal_metadata OWNER TO mac;

--
-- Name: bounce_reports; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.bounce_reports (
    id bigint NOT NULL,
    email public.citext NOT NULL,
    bounce_type character varying NOT NULL,
    created_at timestamp(6) without time zone NOT NULL,
    updated_at timestamp(6) without time zone NOT NULL
);


ALTER TABLE public.bounce_reports OWNER TO mac;

--
-- Name: bounce_reports_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.bounce_reports_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.bounce_reports_id_seq OWNER TO mac;

--
-- Name: bounce_reports_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.bounce_reports_id_seq OWNED BY public.bounce_reports.id;


--
-- Name: certificates; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.certificates (
    id bigint NOT NULL,
    course_id bigint NOT NULL,
    qr_corner character varying NOT NULL,
    qr_scale integer NOT NULL,
    name_offset_top integer NOT NULL,
    font_size integer NOT NULL,
    margin integer NOT NULL,
    active boolean DEFAULT false NOT NULL,
    created_at timestamp(6) without time zone NOT NULL,
    updated_at timestamp(6) without time zone NOT NULL
);


ALTER TABLE public.certificates OWNER TO mac;

--
-- Name: certificates_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.certificates_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.certificates_id_seq OWNER TO mac;

--
-- Name: certificates_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.certificates_id_seq OWNED BY public.certificates.id;


--
-- Name: coach_notes; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.coach_notes (
    id bigint NOT NULL,
    author_id bigint,
    student_id bigint,
    note text,
    created_at timestamp(6) without time zone NOT NULL,
    updated_at timestamp(6) without time zone NOT NULL,
    archived_at timestamp without time zone
);


ALTER TABLE public.coach_notes OWNER TO mac;

--
-- Name: coach_notes_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.coach_notes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.coach_notes_id_seq OWNER TO mac;

--
-- Name: coach_notes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.coach_notes_id_seq OWNED BY public.coach_notes.id;


--
-- Name: colleges; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.colleges (
    id integer NOT NULL,
    name character varying,
    also_known_as character varying,
    city character varying,
    state_id integer,
    established_year character varying,
    website character varying,
    contact_numbers character varying,
    university_id integer
);


ALTER TABLE public.colleges OWNER TO mac;

--
-- Name: colleges_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.colleges_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.colleges_id_seq OWNER TO mac;

--
-- Name: colleges_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.colleges_id_seq OWNED BY public.colleges.id;


--
-- Name: communities; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.communities (
    id bigint NOT NULL,
    name character varying,
    target_linkable boolean DEFAULT false,
    school_id bigint,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.communities OWNER TO mac;

--
-- Name: communities_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.communities_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.communities_id_seq OWNER TO mac;

--
-- Name: communities_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.communities_id_seq OWNED BY public.communities.id;


--
-- Name: community_course_connections; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.community_course_connections (
    id bigint NOT NULL,
    community_id bigint,
    course_id bigint
);


ALTER TABLE public.community_course_connections OWNER TO mac;

--
-- Name: community_course_connections_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.community_course_connections_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.community_course_connections_id_seq OWNER TO mac;

--
-- Name: community_course_connections_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.community_course_connections_id_seq OWNED BY public.community_course_connections.id;


--
-- Name: connect_requests; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.connect_requests (
    id integer NOT NULL,
    connect_slot_id integer,
    startup_id integer,
    questions text,
    status character varying,
    meeting_link character varying,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    confirmed_at timestamp without time zone,
    feedback_mails_sent_at timestamp without time zone,
    rating_for_faculty integer,
    rating_for_team integer,
    comment_for_faculty text,
    comment_for_team text
);


ALTER TABLE public.connect_requests OWNER TO mac;

--
-- Name: connect_requests_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.connect_requests_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.connect_requests_id_seq OWNER TO mac;

--
-- Name: connect_requests_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.connect_requests_id_seq OWNED BY public.connect_requests.id;


--
-- Name: connect_slots; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.connect_slots (
    id integer NOT NULL,
    faculty_id integer,
    slot_at timestamp without time zone,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.connect_slots OWNER TO mac;

--
-- Name: connect_slots_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.connect_slots_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.connect_slots_id_seq OWNER TO mac;

--
-- Name: connect_slots_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.connect_slots_id_seq OWNED BY public.connect_slots.id;


--
-- Name: content_blocks; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.content_blocks (
    id bigint NOT NULL,
    block_type character varying,
    content json,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    sort_index integer DEFAULT 0 NOT NULL,
    target_version_id bigint
);


ALTER TABLE public.content_blocks OWNER TO mac;

--
-- Name: content_blocks_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.content_blocks_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.content_blocks_id_seq OWNER TO mac;

--
-- Name: content_blocks_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.content_blocks_id_seq OWNED BY public.content_blocks.id;


--
-- Name: course_authors; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.course_authors (
    id bigint NOT NULL,
    user_id bigint,
    course_id bigint,
    exited boolean,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.course_authors OWNER TO mac;

--
-- Name: course_authors_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.course_authors_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.course_authors_id_seq OWNER TO mac;

--
-- Name: course_authors_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.course_authors_id_seq OWNED BY public.course_authors.id;


--
-- Name: course_exports; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.course_exports (
    id bigint NOT NULL,
    user_id bigint,
    course_id bigint,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    reviewed_only boolean DEFAULT false,
    json_data text,
    export_type character varying
);


ALTER TABLE public.course_exports OWNER TO mac;

--
-- Name: course_exports_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.course_exports_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.course_exports_id_seq OWNER TO mac;

--
-- Name: course_exports_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.course_exports_id_seq OWNED BY public.course_exports.id;


--
-- Name: courses; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.courses (
    id bigint NOT NULL,
    name character varying,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    school_id bigint,
    ends_at timestamp without time zone,
    description character varying,
    enable_leaderboard boolean DEFAULT false,
    public_signup boolean DEFAULT false,
    about text,
    featured boolean DEFAULT true,
    can_connect boolean DEFAULT true,
    progression_behavior character varying NOT NULL,
    progression_limit integer
);


ALTER TABLE public.courses OWNER TO mac;

--
-- Name: courses_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.courses_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.courses_id_seq OWNER TO mac;

--
-- Name: courses_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.courses_id_seq OWNED BY public.courses.id;


--
-- Name: data_migrations; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.data_migrations (
    version character varying NOT NULL
);


ALTER TABLE public.data_migrations OWNER TO mac;

--
-- Name: delayed_jobs; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.delayed_jobs (
    id integer NOT NULL,
    priority integer DEFAULT 0 NOT NULL,
    attempts integer DEFAULT 0 NOT NULL,
    handler text NOT NULL,
    last_error text,
    run_at timestamp without time zone,
    locked_at timestamp without time zone,
    failed_at timestamp without time zone,
    locked_by character varying,
    queue character varying,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


ALTER TABLE public.delayed_jobs OWNER TO mac;

--
-- Name: delayed_jobs_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.delayed_jobs_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.delayed_jobs_id_seq OWNER TO mac;

--
-- Name: delayed_jobs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.delayed_jobs_id_seq OWNED BY public.delayed_jobs.id;


--
-- Name: domains; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.domains (
    id bigint NOT NULL,
    school_id bigint,
    fqdn character varying,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    "primary" boolean DEFAULT false
);


ALTER TABLE public.domains OWNER TO mac;

--
-- Name: domains_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.domains_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.domains_id_seq OWNER TO mac;

--
-- Name: domains_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.domains_id_seq OWNED BY public.domains.id;


--
-- Name: evaluation_criteria; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.evaluation_criteria (
    id bigint NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    name character varying,
    course_id bigint,
    max_grade integer,
    pass_grade integer,
    grade_labels jsonb
);


ALTER TABLE public.evaluation_criteria OWNER TO mac;

--
-- Name: evaluation_criteria_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.evaluation_criteria_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.evaluation_criteria_id_seq OWNER TO mac;

--
-- Name: evaluation_criteria_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.evaluation_criteria_id_seq OWNED BY public.evaluation_criteria.id;


--
-- Name: faculty; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.faculty (
    id integer NOT NULL,
    category character varying,
    sort_index integer,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    token character varying,
    self_service boolean,
    current_commitment character varying,
    commitment character varying,
    compensation character varying,
    slack_username character varying,
    slack_user_id character varying,
    user_id bigint,
    public boolean DEFAULT false,
    connect_link character varying,
    notify_for_submission boolean DEFAULT false,
    exited boolean DEFAULT false
);


ALTER TABLE public.faculty OWNER TO mac;

--
-- Name: faculty_course_enrollments; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.faculty_course_enrollments (
    id bigint NOT NULL,
    faculty_id bigint,
    course_id bigint,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.faculty_course_enrollments OWNER TO mac;

--
-- Name: faculty_course_enrollments_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.faculty_course_enrollments_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.faculty_course_enrollments_id_seq OWNER TO mac;

--
-- Name: faculty_course_enrollments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.faculty_course_enrollments_id_seq OWNED BY public.faculty_course_enrollments.id;


--
-- Name: faculty_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.faculty_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.faculty_id_seq OWNER TO mac;

--
-- Name: faculty_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.faculty_id_seq OWNED BY public.faculty.id;


--
-- Name: faculty_startup_enrollments; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.faculty_startup_enrollments (
    id bigint NOT NULL,
    faculty_id bigint,
    startup_id bigint,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.faculty_startup_enrollments OWNER TO mac;

--
-- Name: faculty_startup_enrollments_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.faculty_startup_enrollments_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.faculty_startup_enrollments_id_seq OWNER TO mac;

--
-- Name: faculty_startup_enrollments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.faculty_startup_enrollments_id_seq OWNED BY public.faculty_startup_enrollments.id;


--
-- Name: features; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.features (
    id integer NOT NULL,
    key character varying,
    value character varying,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


ALTER TABLE public.features OWNER TO mac;

--
-- Name: features_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.features_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.features_id_seq OWNER TO mac;

--
-- Name: features_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.features_id_seq OWNED BY public.features.id;


--
-- Name: founders; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.founders (
    id integer NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    startup_id integer,
    auth_token character varying,
    slack_username character varying,
    university_id integer,
    roles character varying,
    slack_user_id character varying,
    user_id integer,
    college_id integer,
    dashboard_toured boolean,
    college_text character varying,
    resume_file_id integer,
    slack_access_token character varying,
    excluded_from_leaderboard boolean DEFAULT false
);


ALTER TABLE public.founders OWNER TO mac;

--
-- Name: founders_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.founders_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.founders_id_seq OWNER TO mac;

--
-- Name: founders_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.founders_id_seq OWNED BY public.founders.id;


--
-- Name: issued_certificates; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.issued_certificates (
    id bigint NOT NULL,
    certificate_id bigint NOT NULL,
    user_id bigint NOT NULL,
    name character varying NOT NULL,
    serial_number public.citext NOT NULL,
    created_at timestamp(6) without time zone NOT NULL,
    updated_at timestamp(6) without time zone NOT NULL
);


ALTER TABLE public.issued_certificates OWNER TO mac;

--
-- Name: issued_certificates_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.issued_certificates_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.issued_certificates_id_seq OWNER TO mac;

--
-- Name: issued_certificates_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.issued_certificates_id_seq OWNED BY public.issued_certificates.id;


--
-- Name: leaderboard_entries; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.leaderboard_entries (
    id bigint NOT NULL,
    founder_id bigint,
    period_from timestamp without time zone NOT NULL,
    period_to timestamp without time zone NOT NULL,
    score integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.leaderboard_entries OWNER TO mac;

--
-- Name: leaderboard_entries_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.leaderboard_entries_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.leaderboard_entries_id_seq OWNER TO mac;

--
-- Name: leaderboard_entries_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.leaderboard_entries_id_seq OWNED BY public.leaderboard_entries.id;


--
-- Name: levels; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.levels (
    id integer NOT NULL,
    name character varying,
    description text,
    number integer,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    unlock_on date,
    course_id bigint
);


ALTER TABLE public.levels OWNER TO mac;

--
-- Name: levels_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.levels_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.levels_id_seq OWNER TO mac;

--
-- Name: levels_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.levels_id_seq OWNED BY public.levels.id;


--
-- Name: markdown_attachments; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.markdown_attachments (
    id bigint NOT NULL,
    token character varying,
    last_accessed_at timestamp without time zone,
    user_id bigint,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.markdown_attachments OWNER TO mac;

--
-- Name: markdown_attachments_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.markdown_attachments_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.markdown_attachments_id_seq OWNER TO mac;

--
-- Name: markdown_attachments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.markdown_attachments_id_seq OWNED BY public.markdown_attachments.id;


--
-- Name: platform_feedback; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.platform_feedback (
    id integer NOT NULL,
    feedback_type character varying,
    description text,
    promoter_score integer,
    founder_id integer,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    notes text
);


ALTER TABLE public.platform_feedback OWNER TO mac;

--
-- Name: platform_feedback_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.platform_feedback_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.platform_feedback_id_seq OWNER TO mac;

--
-- Name: platform_feedback_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.platform_feedback_id_seq OWNED BY public.platform_feedback.id;


--
-- Name: post_likes; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.post_likes (
    id bigint NOT NULL,
    post_id bigint,
    user_id bigint,
    created_at timestamp(6) without time zone NOT NULL,
    updated_at timestamp(6) without time zone NOT NULL
);


ALTER TABLE public.post_likes OWNER TO mac;

--
-- Name: post_likes_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.post_likes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.post_likes_id_seq OWNER TO mac;

--
-- Name: post_likes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.post_likes_id_seq OWNED BY public.post_likes.id;


--
-- Name: posts; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.posts (
    id bigint NOT NULL,
    topic_id bigint,
    creator_id bigint,
    editor_id bigint,
    archiver_id bigint,
    archived_at timestamp without time zone,
    reply_to_post_id bigint,
    post_number integer NOT NULL,
    body text,
    solution boolean DEFAULT false,
    created_at timestamp(6) without time zone NOT NULL,
    updated_at timestamp(6) without time zone NOT NULL
);


ALTER TABLE public.posts OWNER TO mac;

--
-- Name: posts_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.posts_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.posts_id_seq OWNER TO mac;

--
-- Name: posts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.posts_id_seq OWNED BY public.posts.id;


--
-- Name: prospective_applicants; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.prospective_applicants (
    id integer NOT NULL,
    name character varying,
    email character varying,
    phone character varying,
    college_id integer,
    college_text character varying,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.prospective_applicants OWNER TO mac;

--
-- Name: prospective_applicants_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.prospective_applicants_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.prospective_applicants_id_seq OWNER TO mac;

--
-- Name: prospective_applicants_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.prospective_applicants_id_seq OWNED BY public.prospective_applicants.id;


--
-- Name: public_slack_messages; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.public_slack_messages (
    id integer NOT NULL,
    body text,
    slack_username character varying,
    founder_id integer,
    channel character varying,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    "timestamp" character varying,
    reaction_to_id integer
);


ALTER TABLE public.public_slack_messages OWNER TO mac;

--
-- Name: public_slack_messages_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.public_slack_messages_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.public_slack_messages_id_seq OWNER TO mac;

--
-- Name: public_slack_messages_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.public_slack_messages_id_seq OWNED BY public.public_slack_messages.id;


--
-- Name: quiz_questions; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.quiz_questions (
    id bigint NOT NULL,
    question text,
    description text,
    quiz_id bigint,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    correct_answer_id bigint
);


ALTER TABLE public.quiz_questions OWNER TO mac;

--
-- Name: quiz_questions_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.quiz_questions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.quiz_questions_id_seq OWNER TO mac;

--
-- Name: quiz_questions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.quiz_questions_id_seq OWNED BY public.quiz_questions.id;


--
-- Name: quizzes; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.quizzes (
    id bigint NOT NULL,
    title character varying,
    target_id bigint,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.quizzes OWNER TO mac;

--
-- Name: quizzes_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.quizzes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.quizzes_id_seq OWNER TO mac;

--
-- Name: quizzes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.quizzes_id_seq OWNED BY public.quizzes.id;


--
-- Name: resource_versions; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.resource_versions (
    id bigint NOT NULL,
    value jsonb,
    versionable_type character varying,
    versionable_id bigint,
    created_at timestamp(6) without time zone NOT NULL,
    updated_at timestamp(6) without time zone NOT NULL
);


ALTER TABLE public.resource_versions OWNER TO mac;

--
-- Name: resource_versions_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.resource_versions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.resource_versions_id_seq OWNER TO mac;

--
-- Name: resource_versions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.resource_versions_id_seq OWNED BY public.resource_versions.id;


--
-- Name: resources; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.resources (
    id integer NOT NULL,
    title character varying,
    description text,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    downloads integer DEFAULT 0,
    slug character varying,
    video_embed text,
    link character varying,
    archived boolean DEFAULT false,
    public boolean DEFAULT false,
    school_id bigint
);


ALTER TABLE public.resources OWNER TO mac;

--
-- Name: resources_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.resources_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.resources_id_seq OWNER TO mac;

--
-- Name: resources_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.resources_id_seq OWNED BY public.resources.id;


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.schema_migrations (
    version character varying NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO mac;

--
-- Name: school_admins; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.school_admins (
    id bigint NOT NULL,
    user_id bigint,
    school_id bigint,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.school_admins OWNER TO mac;

--
-- Name: school_admins_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.school_admins_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.school_admins_id_seq OWNER TO mac;

--
-- Name: school_admins_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.school_admins_id_seq OWNED BY public.school_admins.id;


--
-- Name: school_links; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.school_links (
    id bigint NOT NULL,
    school_id bigint,
    title character varying,
    url character varying,
    kind character varying,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.school_links OWNER TO mac;

--
-- Name: school_links_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.school_links_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.school_links_id_seq OWNER TO mac;

--
-- Name: school_links_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.school_links_id_seq OWNED BY public.school_links.id;


--
-- Name: school_strings; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.school_strings (
    id bigint NOT NULL,
    school_id bigint,
    key character varying,
    value text,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.school_strings OWNER TO mac;

--
-- Name: school_strings_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.school_strings_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.school_strings_id_seq OWNER TO mac;

--
-- Name: school_strings_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.school_strings_id_seq OWNED BY public.school_strings.id;


--
-- Name: schools; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.schools (
    id bigint NOT NULL,
    name character varying,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    about text
);


ALTER TABLE public.schools OWNER TO mac;

--
-- Name: schools_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.schools_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.schools_id_seq OWNER TO mac;

--
-- Name: schools_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.schools_id_seq OWNED BY public.schools.id;


--
-- Name: shortened_urls; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.shortened_urls (
    id integer NOT NULL,
    owner_id integer,
    owner_type character varying(20),
    url text NOT NULL,
    unique_key character varying(100) NOT NULL,
    use_count integer DEFAULT 0 NOT NULL,
    expires_at timestamp without time zone,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


ALTER TABLE public.shortened_urls OWNER TO mac;

--
-- Name: shortened_urls_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.shortened_urls_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.shortened_urls_id_seq OWNER TO mac;

--
-- Name: shortened_urls_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.shortened_urls_id_seq OWNED BY public.shortened_urls.id;


--
-- Name: startup_feedback; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.startup_feedback (
    id integer NOT NULL,
    feedback text,
    reference_url character varying,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    startup_id integer,
    sent_at timestamp without time zone,
    faculty_id integer,
    activity_type character varying,
    attachment character varying,
    timeline_event_id integer
);


ALTER TABLE public.startup_feedback OWNER TO mac;

--
-- Name: startup_feedback_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.startup_feedback_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.startup_feedback_id_seq OWNER TO mac;

--
-- Name: startup_feedback_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.startup_feedback_id_seq OWNED BY public.startup_feedback.id;


--
-- Name: startups; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.startups (
    id integer NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    name character varying,
    slug character varying,
    level_id integer,
    access_ends_at timestamp without time zone,
    dropped_out_at timestamp without time zone
);


ALTER TABLE public.startups OWNER TO mac;

--
-- Name: startups_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.startups_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.startups_id_seq OWNER TO mac;

--
-- Name: startups_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.startups_id_seq OWNED BY public.startups.id;


--
-- Name: states; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.states (
    id integer NOT NULL,
    name character varying,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.states OWNER TO mac;

--
-- Name: states_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.states_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.states_id_seq OWNER TO mac;

--
-- Name: states_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.states_id_seq OWNED BY public.states.id;


--
-- Name: taggings; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.taggings (
    id integer NOT NULL,
    tag_id integer,
    taggable_id integer,
    taggable_type character varying,
    tagger_id integer,
    tagger_type character varying,
    context character varying(128),
    created_at timestamp without time zone
);


ALTER TABLE public.taggings OWNER TO mac;

--
-- Name: taggings_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.taggings_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.taggings_id_seq OWNER TO mac;

--
-- Name: taggings_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.taggings_id_seq OWNED BY public.taggings.id;


--
-- Name: tags; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.tags (
    id integer NOT NULL,
    name character varying,
    taggings_count integer DEFAULT 0
);


ALTER TABLE public.tags OWNER TO mac;

--
-- Name: tags_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.tags_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tags_id_seq OWNER TO mac;

--
-- Name: tags_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.tags_id_seq OWNED BY public.tags.id;


--
-- Name: target_evaluation_criteria; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.target_evaluation_criteria (
    id bigint NOT NULL,
    target_id bigint,
    evaluation_criterion_id bigint,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.target_evaluation_criteria OWNER TO mac;

--
-- Name: target_evaluation_criteria_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.target_evaluation_criteria_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.target_evaluation_criteria_id_seq OWNER TO mac;

--
-- Name: target_evaluation_criteria_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.target_evaluation_criteria_id_seq OWNED BY public.target_evaluation_criteria.id;


--
-- Name: target_groups; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.target_groups (
    id integer NOT NULL,
    name character varying,
    description text,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    sort_index integer,
    milestone boolean,
    level_id integer,
    archived boolean DEFAULT false
);


ALTER TABLE public.target_groups OWNER TO mac;

--
-- Name: target_groups_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.target_groups_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.target_groups_id_seq OWNER TO mac;

--
-- Name: target_groups_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.target_groups_id_seq OWNED BY public.target_groups.id;


--
-- Name: target_prerequisites; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.target_prerequisites (
    id integer NOT NULL,
    target_id integer,
    prerequisite_target_id integer
);


ALTER TABLE public.target_prerequisites OWNER TO mac;

--
-- Name: target_prerequisites_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.target_prerequisites_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.target_prerequisites_id_seq OWNER TO mac;

--
-- Name: target_prerequisites_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.target_prerequisites_id_seq OWNED BY public.target_prerequisites.id;


--
-- Name: target_resources; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.target_resources (
    id bigint NOT NULL,
    target_id bigint NOT NULL,
    resource_id bigint NOT NULL
);


ALTER TABLE public.target_resources OWNER TO mac;

--
-- Name: target_resources_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.target_resources_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.target_resources_id_seq OWNER TO mac;

--
-- Name: target_resources_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.target_resources_id_seq OWNED BY public.target_resources.id;


--
-- Name: target_versions; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.target_versions (
    id bigint NOT NULL,
    target_id bigint,
    created_at timestamp(6) without time zone NOT NULL,
    updated_at timestamp(6) without time zone NOT NULL
);


ALTER TABLE public.target_versions OWNER TO mac;

--
-- Name: target_versions_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.target_versions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.target_versions_id_seq OWNER TO mac;

--
-- Name: target_versions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.target_versions_id_seq OWNED BY public.target_versions.id;


--
-- Name: targets; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.targets (
    id integer NOT NULL,
    role character varying,
    title character varying,
    description text,
    completion_instructions character varying,
    resource_url character varying,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    slideshow_embed text,
    faculty_id integer,
    days_to_complete integer,
    target_action_type character varying,
    target_group_id integer,
    sort_index integer DEFAULT 999,
    session_at timestamp without time zone,
    video_embed text,
    last_session_at timestamp without time zone,
    link_to_complete character varying,
    archived boolean DEFAULT false,
    youtube_video_id character varying,
    google_calendar_event_id character varying,
    feedback_asked_at timestamp without time zone,
    slack_reminders_sent_at timestamp without time zone,
    call_to_action character varying,
    rubric_description text,
    resubmittable boolean DEFAULT true,
    visibility character varying,
    review_checklist jsonb DEFAULT '[]'::jsonb,
    checklist jsonb DEFAULT '[]'::jsonb
);


ALTER TABLE public.targets OWNER TO mac;

--
-- Name: targets_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.targets_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.targets_id_seq OWNER TO mac;

--
-- Name: targets_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.targets_id_seq OWNED BY public.targets.id;


--
-- Name: text_versions; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.text_versions (
    id bigint NOT NULL,
    value text,
    versionable_type character varying,
    versionable_id bigint,
    user_id bigint,
    edited_at timestamp without time zone,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.text_versions OWNER TO mac;

--
-- Name: text_versions_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.text_versions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.text_versions_id_seq OWNER TO mac;

--
-- Name: text_versions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.text_versions_id_seq OWNED BY public.text_versions.id;


--
-- Name: timeline_event_files; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.timeline_event_files (
    id integer NOT NULL,
    timeline_event_id integer,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    title character varying
);


ALTER TABLE public.timeline_event_files OWNER TO mac;

--
-- Name: timeline_event_files_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.timeline_event_files_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.timeline_event_files_id_seq OWNER TO mac;

--
-- Name: timeline_event_files_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.timeline_event_files_id_seq OWNED BY public.timeline_event_files.id;


--
-- Name: timeline_event_grades; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.timeline_event_grades (
    id bigint NOT NULL,
    timeline_event_id bigint,
    evaluation_criterion_id bigint,
    grade integer
);


ALTER TABLE public.timeline_event_grades OWNER TO mac;

--
-- Name: timeline_event_grades_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.timeline_event_grades_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.timeline_event_grades_id_seq OWNER TO mac;

--
-- Name: timeline_event_grades_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.timeline_event_grades_id_seq OWNED BY public.timeline_event_grades.id;


--
-- Name: timeline_event_owners; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.timeline_event_owners (
    id bigint NOT NULL,
    timeline_event_id bigint,
    founder_id bigint,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.timeline_event_owners OWNER TO mac;

--
-- Name: timeline_event_owners_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.timeline_event_owners_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.timeline_event_owners_id_seq OWNER TO mac;

--
-- Name: timeline_event_owners_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.timeline_event_owners_id_seq OWNED BY public.timeline_event_owners.id;


--
-- Name: timeline_events; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.timeline_events (
    id integer NOT NULL,
    image character varying,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    improved_timeline_event_id integer,
    target_id integer,
    score numeric(2,1),
    evaluator_id integer,
    passed_at timestamp without time zone,
    latest boolean,
    quiz_score character varying,
    evaluated_at timestamp without time zone,
    checklist jsonb DEFAULT '[]'::jsonb
);


ALTER TABLE public.timeline_events OWNER TO mac;

--
-- Name: timeline_events_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.timeline_events_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.timeline_events_id_seq OWNER TO mac;

--
-- Name: timeline_events_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.timeline_events_id_seq OWNED BY public.timeline_events.id;


--
-- Name: topics; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.topics (
    id bigint NOT NULL,
    community_id bigint,
    target_id bigint,
    last_activity_at timestamp without time zone,
    archived boolean DEFAULT false NOT NULL,
    title character varying,
    created_at timestamp(6) without time zone NOT NULL,
    updated_at timestamp(6) without time zone NOT NULL
);


ALTER TABLE public.topics OWNER TO mac;

--
-- Name: topics_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.topics_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.topics_id_seq OWNER TO mac;

--
-- Name: topics_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.topics_id_seq OWNED BY public.topics.id;


--
-- Name: universities; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.universities (
    id integer NOT NULL,
    name character varying,
    state_id integer,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.universities OWNER TO mac;

--
-- Name: universities_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.universities_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.universities_id_seq OWNER TO mac;

--
-- Name: universities_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.universities_id_seq OWNED BY public.universities.id;


--
-- Name: user_activities; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.user_activities (
    id integer NOT NULL,
    user_id integer,
    activity_type character varying,
    metadata json,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.user_activities OWNER TO mac;

--
-- Name: user_activities_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.user_activities_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_activities_id_seq OWNER TO mac;

--
-- Name: user_activities_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.user_activities_id_seq OWNED BY public.user_activities.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.users (
    id integer NOT NULL,
    email character varying,
    login_token character varying,
    remember_created_at timestamp without time zone,
    sign_in_count integer DEFAULT 0,
    current_sign_in_at timestamp without time zone,
    last_sign_in_at timestamp without time zone,
    current_sign_in_ip character varying,
    last_sign_in_ip character varying,
    encrypted_password character varying DEFAULT ''::character varying NOT NULL,
    remember_token character varying,
    sign_out_at_next_request boolean,
    confirmed_at timestamp without time zone,
    login_mail_sent_at timestamp without time zone,
    name character varying,
    phone character varying,
    communication_address character varying,
    title character varying,
    key_skills character varying,
    about text,
    resume_url character varying,
    blog_url character varying,
    personal_website_url character varying,
    linkedin_url character varying,
    twitter_url character varying,
    facebook_url character varying,
    angel_co_url character varying,
    github_url character varying,
    behance_url character varying,
    skype_id character varying,
    school_id bigint,
    preferences jsonb DEFAULT '{"daily_digest": true}'::jsonb NOT NULL,
    reset_password_token character varying,
    reset_password_sent_at timestamp without time zone,
    affiliation character varying,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    time_zone character varying DEFAULT 'Asia/Kolkata'::character varying NOT NULL
);


ALTER TABLE public.users OWNER TO mac;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: mac
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO mac;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mac
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: visits; Type: TABLE; Schema: public; Owner: mac
--

CREATE TABLE public.visits (
    id uuid NOT NULL,
    visitor_id uuid,
    ip character varying,
    user_agent text,
    referrer text,
    landing_page text,
    user_id integer,
    user_type character varying,
    referring_domain character varying,
    search_keyword character varying,
    browser character varying,
    os character varying,
    device_type character varying,
    screen_height integer,
    screen_width integer,
    country character varying,
    region character varying,
    city character varying,
    postal_code character varying,
    latitude numeric,
    longitude numeric,
    utm_source character varying,
    utm_medium character varying,
    utm_term character varying,
    utm_content character varying,
    utm_campaign character varying,
    started_at timestamp without time zone
);


ALTER TABLE public.visits OWNER TO mac;

--
-- Name: active_admin_comments id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.active_admin_comments ALTER COLUMN id SET DEFAULT nextval('public.active_admin_comments_id_seq'::regclass);


--
-- Name: active_storage_attachments id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.active_storage_attachments ALTER COLUMN id SET DEFAULT nextval('public.active_storage_attachments_id_seq'::regclass);


--
-- Name: active_storage_blobs id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.active_storage_blobs ALTER COLUMN id SET DEFAULT nextval('public.active_storage_blobs_id_seq'::regclass);


--
-- Name: admin_users id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.admin_users ALTER COLUMN id SET DEFAULT nextval('public.admin_users_id_seq'::regclass);


--
-- Name: answer_options id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.answer_options ALTER COLUMN id SET DEFAULT nextval('public.answer_options_id_seq'::regclass);


--
-- Name: applicants id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.applicants ALTER COLUMN id SET DEFAULT nextval('public.applicants_id_seq'::regclass);


--
-- Name: bounce_reports id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.bounce_reports ALTER COLUMN id SET DEFAULT nextval('public.bounce_reports_id_seq'::regclass);


--
-- Name: certificates id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.certificates ALTER COLUMN id SET DEFAULT nextval('public.certificates_id_seq'::regclass);


--
-- Name: coach_notes id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.coach_notes ALTER COLUMN id SET DEFAULT nextval('public.coach_notes_id_seq'::regclass);


--
-- Name: colleges id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.colleges ALTER COLUMN id SET DEFAULT nextval('public.colleges_id_seq'::regclass);


--
-- Name: communities id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.communities ALTER COLUMN id SET DEFAULT nextval('public.communities_id_seq'::regclass);


--
-- Name: community_course_connections id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.community_course_connections ALTER COLUMN id SET DEFAULT nextval('public.community_course_connections_id_seq'::regclass);


--
-- Name: connect_requests id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.connect_requests ALTER COLUMN id SET DEFAULT nextval('public.connect_requests_id_seq'::regclass);


--
-- Name: connect_slots id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.connect_slots ALTER COLUMN id SET DEFAULT nextval('public.connect_slots_id_seq'::regclass);


--
-- Name: content_blocks id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.content_blocks ALTER COLUMN id SET DEFAULT nextval('public.content_blocks_id_seq'::regclass);


--
-- Name: course_authors id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.course_authors ALTER COLUMN id SET DEFAULT nextval('public.course_authors_id_seq'::regclass);


--
-- Name: course_exports id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.course_exports ALTER COLUMN id SET DEFAULT nextval('public.course_exports_id_seq'::regclass);


--
-- Name: courses id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.courses ALTER COLUMN id SET DEFAULT nextval('public.courses_id_seq'::regclass);


--
-- Name: delayed_jobs id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.delayed_jobs ALTER COLUMN id SET DEFAULT nextval('public.delayed_jobs_id_seq'::regclass);


--
-- Name: domains id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.domains ALTER COLUMN id SET DEFAULT nextval('public.domains_id_seq'::regclass);


--
-- Name: evaluation_criteria id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.evaluation_criteria ALTER COLUMN id SET DEFAULT nextval('public.evaluation_criteria_id_seq'::regclass);


--
-- Name: faculty id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.faculty ALTER COLUMN id SET DEFAULT nextval('public.faculty_id_seq'::regclass);


--
-- Name: faculty_course_enrollments id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.faculty_course_enrollments ALTER COLUMN id SET DEFAULT nextval('public.faculty_course_enrollments_id_seq'::regclass);


--
-- Name: faculty_startup_enrollments id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.faculty_startup_enrollments ALTER COLUMN id SET DEFAULT nextval('public.faculty_startup_enrollments_id_seq'::regclass);


--
-- Name: features id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.features ALTER COLUMN id SET DEFAULT nextval('public.features_id_seq'::regclass);


--
-- Name: founders id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.founders ALTER COLUMN id SET DEFAULT nextval('public.founders_id_seq'::regclass);


--
-- Name: issued_certificates id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.issued_certificates ALTER COLUMN id SET DEFAULT nextval('public.issued_certificates_id_seq'::regclass);


--
-- Name: leaderboard_entries id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.leaderboard_entries ALTER COLUMN id SET DEFAULT nextval('public.leaderboard_entries_id_seq'::regclass);


--
-- Name: levels id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.levels ALTER COLUMN id SET DEFAULT nextval('public.levels_id_seq'::regclass);


--
-- Name: markdown_attachments id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.markdown_attachments ALTER COLUMN id SET DEFAULT nextval('public.markdown_attachments_id_seq'::regclass);


--
-- Name: platform_feedback id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.platform_feedback ALTER COLUMN id SET DEFAULT nextval('public.platform_feedback_id_seq'::regclass);


--
-- Name: post_likes id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.post_likes ALTER COLUMN id SET DEFAULT nextval('public.post_likes_id_seq'::regclass);


--
-- Name: posts id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.posts ALTER COLUMN id SET DEFAULT nextval('public.posts_id_seq'::regclass);


--
-- Name: prospective_applicants id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.prospective_applicants ALTER COLUMN id SET DEFAULT nextval('public.prospective_applicants_id_seq'::regclass);


--
-- Name: public_slack_messages id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.public_slack_messages ALTER COLUMN id SET DEFAULT nextval('public.public_slack_messages_id_seq'::regclass);


--
-- Name: quiz_questions id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.quiz_questions ALTER COLUMN id SET DEFAULT nextval('public.quiz_questions_id_seq'::regclass);


--
-- Name: quizzes id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.quizzes ALTER COLUMN id SET DEFAULT nextval('public.quizzes_id_seq'::regclass);


--
-- Name: resource_versions id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.resource_versions ALTER COLUMN id SET DEFAULT nextval('public.resource_versions_id_seq'::regclass);


--
-- Name: resources id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.resources ALTER COLUMN id SET DEFAULT nextval('public.resources_id_seq'::regclass);


--
-- Name: school_admins id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.school_admins ALTER COLUMN id SET DEFAULT nextval('public.school_admins_id_seq'::regclass);


--
-- Name: school_links id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.school_links ALTER COLUMN id SET DEFAULT nextval('public.school_links_id_seq'::regclass);


--
-- Name: school_strings id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.school_strings ALTER COLUMN id SET DEFAULT nextval('public.school_strings_id_seq'::regclass);


--
-- Name: schools id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.schools ALTER COLUMN id SET DEFAULT nextval('public.schools_id_seq'::regclass);


--
-- Name: shortened_urls id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.shortened_urls ALTER COLUMN id SET DEFAULT nextval('public.shortened_urls_id_seq'::regclass);


--
-- Name: startup_feedback id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.startup_feedback ALTER COLUMN id SET DEFAULT nextval('public.startup_feedback_id_seq'::regclass);


--
-- Name: startups id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.startups ALTER COLUMN id SET DEFAULT nextval('public.startups_id_seq'::regclass);


--
-- Name: states id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.states ALTER COLUMN id SET DEFAULT nextval('public.states_id_seq'::regclass);


--
-- Name: taggings id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.taggings ALTER COLUMN id SET DEFAULT nextval('public.taggings_id_seq'::regclass);


--
-- Name: tags id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.tags ALTER COLUMN id SET DEFAULT nextval('public.tags_id_seq'::regclass);


--
-- Name: target_evaluation_criteria id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.target_evaluation_criteria ALTER COLUMN id SET DEFAULT nextval('public.target_evaluation_criteria_id_seq'::regclass);


--
-- Name: target_groups id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.target_groups ALTER COLUMN id SET DEFAULT nextval('public.target_groups_id_seq'::regclass);


--
-- Name: target_prerequisites id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.target_prerequisites ALTER COLUMN id SET DEFAULT nextval('public.target_prerequisites_id_seq'::regclass);


--
-- Name: target_resources id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.target_resources ALTER COLUMN id SET DEFAULT nextval('public.target_resources_id_seq'::regclass);


--
-- Name: target_versions id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.target_versions ALTER COLUMN id SET DEFAULT nextval('public.target_versions_id_seq'::regclass);


--
-- Name: targets id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.targets ALTER COLUMN id SET DEFAULT nextval('public.targets_id_seq'::regclass);


--
-- Name: text_versions id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.text_versions ALTER COLUMN id SET DEFAULT nextval('public.text_versions_id_seq'::regclass);


--
-- Name: timeline_event_files id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.timeline_event_files ALTER COLUMN id SET DEFAULT nextval('public.timeline_event_files_id_seq'::regclass);


--
-- Name: timeline_event_grades id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.timeline_event_grades ALTER COLUMN id SET DEFAULT nextval('public.timeline_event_grades_id_seq'::regclass);


--
-- Name: timeline_event_owners id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.timeline_event_owners ALTER COLUMN id SET DEFAULT nextval('public.timeline_event_owners_id_seq'::regclass);


--
-- Name: timeline_events id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.timeline_events ALTER COLUMN id SET DEFAULT nextval('public.timeline_events_id_seq'::regclass);


--
-- Name: topics id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.topics ALTER COLUMN id SET DEFAULT nextval('public.topics_id_seq'::regclass);


--
-- Name: universities id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.universities ALTER COLUMN id SET DEFAULT nextval('public.universities_id_seq'::regclass);


--
-- Name: user_activities id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.user_activities ALTER COLUMN id SET DEFAULT nextval('public.user_activities_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: active_admin_comments active_admin_comments_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.active_admin_comments
    ADD CONSTRAINT active_admin_comments_pkey PRIMARY KEY (id);


--
-- Name: active_storage_attachments active_storage_attachments_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.active_storage_attachments
    ADD CONSTRAINT active_storage_attachments_pkey PRIMARY KEY (id);


--
-- Name: active_storage_blobs active_storage_blobs_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.active_storage_blobs
    ADD CONSTRAINT active_storage_blobs_pkey PRIMARY KEY (id);


--
-- Name: admin_users admin_users_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.admin_users
    ADD CONSTRAINT admin_users_pkey PRIMARY KEY (id);


--
-- Name: ahoy_events ahoy_events_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.ahoy_events
    ADD CONSTRAINT ahoy_events_pkey PRIMARY KEY (id);


--
-- Name: answer_options answer_options_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.answer_options
    ADD CONSTRAINT answer_options_pkey PRIMARY KEY (id);


--
-- Name: applicants applicants_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.applicants
    ADD CONSTRAINT applicants_pkey PRIMARY KEY (id);


--
-- Name: ar_internal_metadata ar_internal_metadata_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.ar_internal_metadata
    ADD CONSTRAINT ar_internal_metadata_pkey PRIMARY KEY (key);


--
-- Name: bounce_reports bounce_reports_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.bounce_reports
    ADD CONSTRAINT bounce_reports_pkey PRIMARY KEY (id);


--
-- Name: certificates certificates_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.certificates
    ADD CONSTRAINT certificates_pkey PRIMARY KEY (id);


--
-- Name: coach_notes coach_notes_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.coach_notes
    ADD CONSTRAINT coach_notes_pkey PRIMARY KEY (id);


--
-- Name: colleges colleges_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.colleges
    ADD CONSTRAINT colleges_pkey PRIMARY KEY (id);


--
-- Name: communities communities_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.communities
    ADD CONSTRAINT communities_pkey PRIMARY KEY (id);


--
-- Name: community_course_connections community_course_connections_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.community_course_connections
    ADD CONSTRAINT community_course_connections_pkey PRIMARY KEY (id);


--
-- Name: connect_requests connect_requests_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.connect_requests
    ADD CONSTRAINT connect_requests_pkey PRIMARY KEY (id);


--
-- Name: connect_slots connect_slots_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.connect_slots
    ADD CONSTRAINT connect_slots_pkey PRIMARY KEY (id);


--
-- Name: content_blocks content_blocks_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.content_blocks
    ADD CONSTRAINT content_blocks_pkey PRIMARY KEY (id);


--
-- Name: course_authors course_authors_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.course_authors
    ADD CONSTRAINT course_authors_pkey PRIMARY KEY (id);


--
-- Name: course_exports course_exports_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.course_exports
    ADD CONSTRAINT course_exports_pkey PRIMARY KEY (id);


--
-- Name: courses courses_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.courses
    ADD CONSTRAINT courses_pkey PRIMARY KEY (id);


--
-- Name: delayed_jobs delayed_jobs_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.delayed_jobs
    ADD CONSTRAINT delayed_jobs_pkey PRIMARY KEY (id);


--
-- Name: domains domains_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.domains
    ADD CONSTRAINT domains_pkey PRIMARY KEY (id);


--
-- Name: evaluation_criteria evaluation_criteria_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.evaluation_criteria
    ADD CONSTRAINT evaluation_criteria_pkey PRIMARY KEY (id);


--
-- Name: faculty_course_enrollments faculty_course_enrollments_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.faculty_course_enrollments
    ADD CONSTRAINT faculty_course_enrollments_pkey PRIMARY KEY (id);


--
-- Name: faculty faculty_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.faculty
    ADD CONSTRAINT faculty_pkey PRIMARY KEY (id);


--
-- Name: faculty_startup_enrollments faculty_startup_enrollments_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.faculty_startup_enrollments
    ADD CONSTRAINT faculty_startup_enrollments_pkey PRIMARY KEY (id);


--
-- Name: features features_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.features
    ADD CONSTRAINT features_pkey PRIMARY KEY (id);


--
-- Name: founders founders_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.founders
    ADD CONSTRAINT founders_pkey PRIMARY KEY (id);


--
-- Name: issued_certificates issued_certificates_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.issued_certificates
    ADD CONSTRAINT issued_certificates_pkey PRIMARY KEY (id);


--
-- Name: leaderboard_entries leaderboard_entries_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.leaderboard_entries
    ADD CONSTRAINT leaderboard_entries_pkey PRIMARY KEY (id);


--
-- Name: levels levels_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.levels
    ADD CONSTRAINT levels_pkey PRIMARY KEY (id);


--
-- Name: markdown_attachments markdown_attachments_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.markdown_attachments
    ADD CONSTRAINT markdown_attachments_pkey PRIMARY KEY (id);


--
-- Name: platform_feedback platform_feedback_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.platform_feedback
    ADD CONSTRAINT platform_feedback_pkey PRIMARY KEY (id);


--
-- Name: post_likes post_likes_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.post_likes
    ADD CONSTRAINT post_likes_pkey PRIMARY KEY (id);


--
-- Name: posts posts_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_pkey PRIMARY KEY (id);


--
-- Name: prospective_applicants prospective_applicants_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.prospective_applicants
    ADD CONSTRAINT prospective_applicants_pkey PRIMARY KEY (id);


--
-- Name: public_slack_messages public_slack_messages_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.public_slack_messages
    ADD CONSTRAINT public_slack_messages_pkey PRIMARY KEY (id);


--
-- Name: quiz_questions quiz_questions_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.quiz_questions
    ADD CONSTRAINT quiz_questions_pkey PRIMARY KEY (id);


--
-- Name: quizzes quizzes_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.quizzes
    ADD CONSTRAINT quizzes_pkey PRIMARY KEY (id);


--
-- Name: resource_versions resource_versions_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.resource_versions
    ADD CONSTRAINT resource_versions_pkey PRIMARY KEY (id);


--
-- Name: resources resources_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.resources
    ADD CONSTRAINT resources_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: school_admins school_admins_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.school_admins
    ADD CONSTRAINT school_admins_pkey PRIMARY KEY (id);


--
-- Name: school_links school_links_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.school_links
    ADD CONSTRAINT school_links_pkey PRIMARY KEY (id);


--
-- Name: school_strings school_strings_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.school_strings
    ADD CONSTRAINT school_strings_pkey PRIMARY KEY (id);


--
-- Name: schools schools_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.schools
    ADD CONSTRAINT schools_pkey PRIMARY KEY (id);


--
-- Name: shortened_urls shortened_urls_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.shortened_urls
    ADD CONSTRAINT shortened_urls_pkey PRIMARY KEY (id);


--
-- Name: startup_feedback startup_feedback_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.startup_feedback
    ADD CONSTRAINT startup_feedback_pkey PRIMARY KEY (id);


--
-- Name: startups startups_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.startups
    ADD CONSTRAINT startups_pkey PRIMARY KEY (id);


--
-- Name: states states_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.states
    ADD CONSTRAINT states_pkey PRIMARY KEY (id);


--
-- Name: taggings taggings_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.taggings
    ADD CONSTRAINT taggings_pkey PRIMARY KEY (id);


--
-- Name: tags tags_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.tags
    ADD CONSTRAINT tags_pkey PRIMARY KEY (id);


--
-- Name: target_evaluation_criteria target_evaluation_criteria_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.target_evaluation_criteria
    ADD CONSTRAINT target_evaluation_criteria_pkey PRIMARY KEY (id);


--
-- Name: target_groups target_groups_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.target_groups
    ADD CONSTRAINT target_groups_pkey PRIMARY KEY (id);


--
-- Name: target_prerequisites target_prerequisites_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.target_prerequisites
    ADD CONSTRAINT target_prerequisites_pkey PRIMARY KEY (id);


--
-- Name: target_resources target_resources_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.target_resources
    ADD CONSTRAINT target_resources_pkey PRIMARY KEY (id);


--
-- Name: target_versions target_versions_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.target_versions
    ADD CONSTRAINT target_versions_pkey PRIMARY KEY (id);


--
-- Name: targets targets_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.targets
    ADD CONSTRAINT targets_pkey PRIMARY KEY (id);


--
-- Name: text_versions text_versions_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.text_versions
    ADD CONSTRAINT text_versions_pkey PRIMARY KEY (id);


--
-- Name: timeline_event_files timeline_event_files_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.timeline_event_files
    ADD CONSTRAINT timeline_event_files_pkey PRIMARY KEY (id);


--
-- Name: timeline_event_grades timeline_event_grades_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.timeline_event_grades
    ADD CONSTRAINT timeline_event_grades_pkey PRIMARY KEY (id);


--
-- Name: timeline_event_owners timeline_event_owners_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.timeline_event_owners
    ADD CONSTRAINT timeline_event_owners_pkey PRIMARY KEY (id);


--
-- Name: timeline_events timeline_events_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.timeline_events
    ADD CONSTRAINT timeline_events_pkey PRIMARY KEY (id);


--
-- Name: topics topics_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.topics
    ADD CONSTRAINT topics_pkey PRIMARY KEY (id);


--
-- Name: universities universities_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.universities
    ADD CONSTRAINT universities_pkey PRIMARY KEY (id);


--
-- Name: user_activities user_activities_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.user_activities
    ADD CONSTRAINT user_activities_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: visits visits_pkey; Type: CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.visits
    ADD CONSTRAINT visits_pkey PRIMARY KEY (id);


--
-- Name: by_timeline_event_criterion; Type: INDEX; Schema: public; Owner: mac
--

CREATE UNIQUE INDEX by_timeline_event_criterion ON public.timeline_event_grades USING btree (timeline_event_id, evaluation_criterion_id);


--
-- Name: delayed_jobs_priority; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX delayed_jobs_priority ON public.delayed_jobs USING btree (priority, run_at);


--
-- Name: index_active_admin_comments_on_author_type_and_author_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_active_admin_comments_on_author_type_and_author_id ON public.active_admin_comments USING btree (author_type, author_id);


--
-- Name: index_active_admin_comments_on_namespace; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_active_admin_comments_on_namespace ON public.active_admin_comments USING btree (namespace);


--
-- Name: index_active_admin_comments_on_resource_type_and_resource_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_active_admin_comments_on_resource_type_and_resource_id ON public.active_admin_comments USING btree (resource_type, resource_id);


--
-- Name: index_active_storage_attachments_on_blob_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_active_storage_attachments_on_blob_id ON public.active_storage_attachments USING btree (blob_id);


--
-- Name: index_active_storage_attachments_uniqueness; Type: INDEX; Schema: public; Owner: mac
--

CREATE UNIQUE INDEX index_active_storage_attachments_uniqueness ON public.active_storage_attachments USING btree (record_type, record_id, name, blob_id);


--
-- Name: index_active_storage_blobs_on_key; Type: INDEX; Schema: public; Owner: mac
--

CREATE UNIQUE INDEX index_active_storage_blobs_on_key ON public.active_storage_blobs USING btree (key);


--
-- Name: index_admin_users_on_user_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_admin_users_on_user_id ON public.admin_users USING btree (user_id);


--
-- Name: index_ahoy_events_on_time; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_ahoy_events_on_time ON public.ahoy_events USING btree ("time");


--
-- Name: index_ahoy_events_on_user_id_and_user_type; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_ahoy_events_on_user_id_and_user_type ON public.ahoy_events USING btree (user_id, user_type);


--
-- Name: index_ahoy_events_on_visit_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_ahoy_events_on_visit_id ON public.ahoy_events USING btree (visit_id);


--
-- Name: index_answer_options_on_quiz_question_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_answer_options_on_quiz_question_id ON public.answer_options USING btree (quiz_question_id);


--
-- Name: index_applicants_on_course_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_applicants_on_course_id ON public.applicants USING btree (course_id);


--
-- Name: index_applicants_on_email_and_course_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE UNIQUE INDEX index_applicants_on_email_and_course_id ON public.applicants USING btree (email, course_id);


--
-- Name: index_applicants_on_login_token; Type: INDEX; Schema: public; Owner: mac
--

CREATE UNIQUE INDEX index_applicants_on_login_token ON public.applicants USING btree (login_token);


--
-- Name: index_bounce_reports_on_email; Type: INDEX; Schema: public; Owner: mac
--

CREATE UNIQUE INDEX index_bounce_reports_on_email ON public.bounce_reports USING btree (email);


--
-- Name: index_certificates_on_course_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_certificates_on_course_id ON public.certificates USING btree (course_id);


--
-- Name: index_coach_notes_on_archived_at; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_coach_notes_on_archived_at ON public.coach_notes USING btree (archived_at);


--
-- Name: index_coach_notes_on_author_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_coach_notes_on_author_id ON public.coach_notes USING btree (author_id);


--
-- Name: index_coach_notes_on_student_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_coach_notes_on_student_id ON public.coach_notes USING btree (student_id);


--
-- Name: index_colleges_on_state_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_colleges_on_state_id ON public.colleges USING btree (state_id);


--
-- Name: index_colleges_on_university_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_colleges_on_university_id ON public.colleges USING btree (university_id);


--
-- Name: index_communities_on_school_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_communities_on_school_id ON public.communities USING btree (school_id);


--
-- Name: index_community_course_connection_on_course_id_and_community_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE UNIQUE INDEX index_community_course_connection_on_course_id_and_community_id ON public.community_course_connections USING btree (course_id, community_id);


--
-- Name: index_community_course_connections_on_community_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_community_course_connections_on_community_id ON public.community_course_connections USING btree (community_id);


--
-- Name: index_connect_requests_on_connect_slot_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_connect_requests_on_connect_slot_id ON public.connect_requests USING btree (connect_slot_id);


--
-- Name: index_connect_requests_on_startup_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_connect_requests_on_startup_id ON public.connect_requests USING btree (startup_id);


--
-- Name: index_connect_slots_on_faculty_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_connect_slots_on_faculty_id ON public.connect_slots USING btree (faculty_id);


--
-- Name: index_content_blocks_on_block_type; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_content_blocks_on_block_type ON public.content_blocks USING btree (block_type);


--
-- Name: index_content_blocks_on_target_version_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_content_blocks_on_target_version_id ON public.content_blocks USING btree (target_version_id);


--
-- Name: index_course_authors_on_course_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_course_authors_on_course_id ON public.course_authors USING btree (course_id);


--
-- Name: index_course_authors_on_user_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_course_authors_on_user_id ON public.course_authors USING btree (user_id);


--
-- Name: index_course_exports_on_course_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_course_exports_on_course_id ON public.course_exports USING btree (course_id);


--
-- Name: index_course_exports_on_user_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_course_exports_on_user_id ON public.course_exports USING btree (user_id);


--
-- Name: index_courses_on_school_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_courses_on_school_id ON public.courses USING btree (school_id);


--
-- Name: index_domains_on_fqdn; Type: INDEX; Schema: public; Owner: mac
--

CREATE UNIQUE INDEX index_domains_on_fqdn ON public.domains USING btree (fqdn);


--
-- Name: index_domains_on_school_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_domains_on_school_id ON public.domains USING btree (school_id);


--
-- Name: index_evaluation_criteria_on_course_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_evaluation_criteria_on_course_id ON public.evaluation_criteria USING btree (course_id);


--
-- Name: index_faculty_course_enrollments_on_course_id_and_faculty_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE UNIQUE INDEX index_faculty_course_enrollments_on_course_id_and_faculty_id ON public.faculty_course_enrollments USING btree (course_id, faculty_id);


--
-- Name: index_faculty_course_enrollments_on_faculty_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_faculty_course_enrollments_on_faculty_id ON public.faculty_course_enrollments USING btree (faculty_id);


--
-- Name: index_faculty_on_category; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_faculty_on_category ON public.faculty USING btree (category);


--
-- Name: index_faculty_on_user_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_faculty_on_user_id ON public.faculty USING btree (user_id);


--
-- Name: index_faculty_startup_enrollments_on_faculty_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_faculty_startup_enrollments_on_faculty_id ON public.faculty_startup_enrollments USING btree (faculty_id);


--
-- Name: index_faculty_startup_enrollments_on_startup_id_and_faculty_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE UNIQUE INDEX index_faculty_startup_enrollments_on_startup_id_and_faculty_id ON public.faculty_startup_enrollments USING btree (startup_id, faculty_id);


--
-- Name: index_founders_on_college_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_founders_on_college_id ON public.founders USING btree (college_id);


--
-- Name: index_founders_on_university_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_founders_on_university_id ON public.founders USING btree (university_id);


--
-- Name: index_founders_on_user_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_founders_on_user_id ON public.founders USING btree (user_id);


--
-- Name: index_issued_certificates_on_certificate_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_issued_certificates_on_certificate_id ON public.issued_certificates USING btree (certificate_id);


--
-- Name: index_issued_certificates_on_serial_number; Type: INDEX; Schema: public; Owner: mac
--

CREATE UNIQUE INDEX index_issued_certificates_on_serial_number ON public.issued_certificates USING btree (serial_number);


--
-- Name: index_issued_certificates_on_user_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_issued_certificates_on_user_id ON public.issued_certificates USING btree (user_id);


--
-- Name: index_leaderboard_entries_on_founder_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_leaderboard_entries_on_founder_id ON public.leaderboard_entries USING btree (founder_id);


--
-- Name: index_leaderboard_entries_on_founder_id_and_period; Type: INDEX; Schema: public; Owner: mac
--

CREATE UNIQUE INDEX index_leaderboard_entries_on_founder_id_and_period ON public.leaderboard_entries USING btree (founder_id, period_from, period_to);


--
-- Name: index_levels_on_course_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_levels_on_course_id ON public.levels USING btree (course_id);


--
-- Name: index_levels_on_number_and_course_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE UNIQUE INDEX index_levels_on_number_and_course_id ON public.levels USING btree (number, course_id);


--
-- Name: index_markdown_attachments_on_user_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_markdown_attachments_on_user_id ON public.markdown_attachments USING btree (user_id);


--
-- Name: index_platform_feedback_on_founder_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_platform_feedback_on_founder_id ON public.platform_feedback USING btree (founder_id);


--
-- Name: index_post_likes_on_post_id_and_user_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE UNIQUE INDEX index_post_likes_on_post_id_and_user_id ON public.post_likes USING btree (post_id, user_id);


--
-- Name: index_post_likes_on_user_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_post_likes_on_user_id ON public.post_likes USING btree (user_id);


--
-- Name: index_posts_on_archiver_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_posts_on_archiver_id ON public.posts USING btree (archiver_id);


--
-- Name: index_posts_on_creator_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_posts_on_creator_id ON public.posts USING btree (creator_id);


--
-- Name: index_posts_on_editor_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_posts_on_editor_id ON public.posts USING btree (editor_id);


--
-- Name: index_posts_on_post_number_and_topic_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE UNIQUE INDEX index_posts_on_post_number_and_topic_id ON public.posts USING btree (post_number, topic_id);


--
-- Name: index_posts_on_reply_to_post_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_posts_on_reply_to_post_id ON public.posts USING btree (reply_to_post_id);


--
-- Name: index_posts_on_topic_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_posts_on_topic_id ON public.posts USING btree (topic_id);


--
-- Name: index_prospective_applicants_on_college_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_prospective_applicants_on_college_id ON public.prospective_applicants USING btree (college_id);


--
-- Name: index_public_slack_messages_on_founder_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_public_slack_messages_on_founder_id ON public.public_slack_messages USING btree (founder_id);


--
-- Name: index_quiz_questions_on_correct_answer_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_quiz_questions_on_correct_answer_id ON public.quiz_questions USING btree (correct_answer_id);


--
-- Name: index_quiz_questions_on_quiz_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_quiz_questions_on_quiz_id ON public.quiz_questions USING btree (quiz_id);


--
-- Name: index_quizzes_on_target_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_quizzes_on_target_id ON public.quizzes USING btree (target_id);


--
-- Name: index_resource_versions_on_versionable_type_and_versionable_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_resource_versions_on_versionable_type_and_versionable_id ON public.resource_versions USING btree (versionable_type, versionable_id);


--
-- Name: index_resources_on_archived; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_resources_on_archived ON public.resources USING btree (archived);


--
-- Name: index_resources_on_school_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_resources_on_school_id ON public.resources USING btree (school_id);


--
-- Name: index_resources_on_slug; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_resources_on_slug ON public.resources USING btree (slug);


--
-- Name: index_school_admins_on_school_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_school_admins_on_school_id ON public.school_admins USING btree (school_id);


--
-- Name: index_school_admins_on_user_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_school_admins_on_user_id ON public.school_admins USING btree (user_id);


--
-- Name: index_school_admins_on_user_id_and_school_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE UNIQUE INDEX index_school_admins_on_user_id_and_school_id ON public.school_admins USING btree (user_id, school_id);


--
-- Name: index_school_links_on_school_id_and_kind; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_school_links_on_school_id_and_kind ON public.school_links USING btree (school_id, kind);


--
-- Name: index_school_strings_on_school_id_and_key; Type: INDEX; Schema: public; Owner: mac
--

CREATE UNIQUE INDEX index_school_strings_on_school_id_and_key ON public.school_strings USING btree (school_id, key);


--
-- Name: index_shortened_urls_on_owner_id_and_owner_type; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_shortened_urls_on_owner_id_and_owner_type ON public.shortened_urls USING btree (owner_id, owner_type);


--
-- Name: index_shortened_urls_on_unique_key; Type: INDEX; Schema: public; Owner: mac
--

CREATE UNIQUE INDEX index_shortened_urls_on_unique_key ON public.shortened_urls USING btree (unique_key);


--
-- Name: index_shortened_urls_on_url; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_shortened_urls_on_url ON public.shortened_urls USING btree (url);


--
-- Name: index_startup_feedback_on_faculty_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_startup_feedback_on_faculty_id ON public.startup_feedback USING btree (faculty_id);


--
-- Name: index_startup_feedback_on_timeline_event_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_startup_feedback_on_timeline_event_id ON public.startup_feedback USING btree (timeline_event_id);


--
-- Name: index_startups_on_level_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_startups_on_level_id ON public.startups USING btree (level_id);


--
-- Name: index_startups_on_slug; Type: INDEX; Schema: public; Owner: mac
--

CREATE UNIQUE INDEX index_startups_on_slug ON public.startups USING btree (slug);


--
-- Name: index_taggings_on_context; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_taggings_on_context ON public.taggings USING btree (context);


--
-- Name: index_taggings_on_tag_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_taggings_on_tag_id ON public.taggings USING btree (tag_id);


--
-- Name: index_taggings_on_taggable_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_taggings_on_taggable_id ON public.taggings USING btree (taggable_id);


--
-- Name: index_taggings_on_taggable_id_and_taggable_type_and_context; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_taggings_on_taggable_id_and_taggable_type_and_context ON public.taggings USING btree (taggable_id, taggable_type, context);


--
-- Name: index_taggings_on_taggable_type; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_taggings_on_taggable_type ON public.taggings USING btree (taggable_type);


--
-- Name: index_taggings_on_tagger_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_taggings_on_tagger_id ON public.taggings USING btree (tagger_id);


--
-- Name: index_taggings_on_tagger_id_and_tagger_type; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_taggings_on_tagger_id_and_tagger_type ON public.taggings USING btree (tagger_id, tagger_type);


--
-- Name: index_tags_on_name; Type: INDEX; Schema: public; Owner: mac
--

CREATE UNIQUE INDEX index_tags_on_name ON public.tags USING btree (name);


--
-- Name: index_target_evaluation_criteria_on_evaluation_criterion_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_target_evaluation_criteria_on_evaluation_criterion_id ON public.target_evaluation_criteria USING btree (evaluation_criterion_id);


--
-- Name: index_target_evaluation_criteria_on_target_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_target_evaluation_criteria_on_target_id ON public.target_evaluation_criteria USING btree (target_id);


--
-- Name: index_target_groups_on_level_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_target_groups_on_level_id ON public.target_groups USING btree (level_id);


--
-- Name: index_target_groups_on_sort_index; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_target_groups_on_sort_index ON public.target_groups USING btree (sort_index);


--
-- Name: index_target_prerequisites_on_prerequisite_target_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_target_prerequisites_on_prerequisite_target_id ON public.target_prerequisites USING btree (prerequisite_target_id);


--
-- Name: index_target_prerequisites_on_target_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_target_prerequisites_on_target_id ON public.target_prerequisites USING btree (target_id);


--
-- Name: index_target_resources_on_resource_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_target_resources_on_resource_id ON public.target_resources USING btree (resource_id);


--
-- Name: index_target_resources_on_target_id_and_resource_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE UNIQUE INDEX index_target_resources_on_target_id_and_resource_id ON public.target_resources USING btree (target_id, resource_id);


--
-- Name: index_target_versions_on_target_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_target_versions_on_target_id ON public.target_versions USING btree (target_id);


--
-- Name: index_targets_on_archived; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_targets_on_archived ON public.targets USING btree (archived);


--
-- Name: index_targets_on_faculty_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_targets_on_faculty_id ON public.targets USING btree (faculty_id);


--
-- Name: index_targets_on_session_at; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_targets_on_session_at ON public.targets USING btree (session_at);


--
-- Name: index_text_versions_on_user_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_text_versions_on_user_id ON public.text_versions USING btree (user_id);


--
-- Name: index_text_versions_on_versionable_type_and_versionable_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_text_versions_on_versionable_type_and_versionable_id ON public.text_versions USING btree (versionable_type, versionable_id);


--
-- Name: index_timeline_event_files_on_timeline_event_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_timeline_event_files_on_timeline_event_id ON public.timeline_event_files USING btree (timeline_event_id);


--
-- Name: index_timeline_event_grades_on_evaluation_criterion_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_timeline_event_grades_on_evaluation_criterion_id ON public.timeline_event_grades USING btree (evaluation_criterion_id);


--
-- Name: index_timeline_event_grades_on_timeline_event_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_timeline_event_grades_on_timeline_event_id ON public.timeline_event_grades USING btree (timeline_event_id);


--
-- Name: index_timeline_event_owners_on_founder_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_timeline_event_owners_on_founder_id ON public.timeline_event_owners USING btree (founder_id);


--
-- Name: index_timeline_event_owners_on_timeline_event_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_timeline_event_owners_on_timeline_event_id ON public.timeline_event_owners USING btree (timeline_event_id);


--
-- Name: index_topics_on_community_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_topics_on_community_id ON public.topics USING btree (community_id);


--
-- Name: index_topics_on_target_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_topics_on_target_id ON public.topics USING btree (target_id);


--
-- Name: index_universities_on_state_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_universities_on_state_id ON public.universities USING btree (state_id);


--
-- Name: index_user_activities_on_user_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_user_activities_on_user_id ON public.user_activities USING btree (user_id);


--
-- Name: index_users_on_email_and_school_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE UNIQUE INDEX index_users_on_email_and_school_id ON public.users USING btree (email, school_id);


--
-- Name: index_users_on_reset_password_token; Type: INDEX; Schema: public; Owner: mac
--

CREATE UNIQUE INDEX index_users_on_reset_password_token ON public.users USING btree (reset_password_token);


--
-- Name: index_users_on_school_id; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_users_on_school_id ON public.users USING btree (school_id);


--
-- Name: index_visits_on_user_id_and_user_type; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX index_visits_on_user_id_and_user_type ON public.visits USING btree (user_id, user_type);


--
-- Name: taggings_idx; Type: INDEX; Schema: public; Owner: mac
--

CREATE UNIQUE INDEX taggings_idx ON public.taggings USING btree (tag_id, taggable_id, taggable_type, context, tagger_id, tagger_type);


--
-- Name: taggings_idy; Type: INDEX; Schema: public; Owner: mac
--

CREATE INDEX taggings_idy ON public.taggings USING btree (taggable_id, taggable_type, tagger_id, context);


--
-- Name: unique_data_migrations; Type: INDEX; Schema: public; Owner: mac
--

CREATE UNIQUE INDEX unique_data_migrations ON public.data_migrations USING btree (version);


--
-- Name: founders fk_rails_0066a6b9d3; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.founders
    ADD CONSTRAINT fk_rails_0066a6b9d3 FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: connect_requests fk_rails_04405fdc41; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.connect_requests
    ADD CONSTRAINT fk_rails_04405fdc41 FOREIGN KEY (startup_id) REFERENCES public.startups(id);


--
-- Name: startup_feedback fk_rails_0441f728b8; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.startup_feedback
    ADD CONSTRAINT fk_rails_0441f728b8 FOREIGN KEY (timeline_event_id) REFERENCES public.timeline_events(id);


--
-- Name: leaderboard_entries fk_rails_04e22c1259; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.leaderboard_entries
    ADD CONSTRAINT fk_rails_04e22c1259 FOREIGN KEY (founder_id) REFERENCES public.founders(id);


--
-- Name: posts fk_rails_08c1eddbe0; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT fk_rails_08c1eddbe0 FOREIGN KEY (reply_to_post_id) REFERENCES public.posts(id);


--
-- Name: answer_options fk_rails_1531478821; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.answer_options
    ADD CONSTRAINT fk_rails_1531478821 FOREIGN KEY (quiz_question_id) REFERENCES public.quiz_questions(id);


--
-- Name: community_course_connections fk_rails_23a3259f0e; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.community_course_connections
    ADD CONSTRAINT fk_rails_23a3259f0e FOREIGN KEY (course_id) REFERENCES public.courses(id);


--
-- Name: markdown_attachments fk_rails_306333d224; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.markdown_attachments
    ADD CONSTRAINT fk_rails_306333d224 FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: topics fk_rails_385afaa68c; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.topics
    ADD CONSTRAINT fk_rails_385afaa68c FOREIGN KEY (community_id) REFERENCES public.communities(id);


--
-- Name: school_admins fk_rails_4a2a343dbc; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.school_admins
    ADD CONSTRAINT fk_rails_4a2a343dbc FOREIGN KEY (school_id) REFERENCES public.schools(id);


--
-- Name: certificates fk_rails_4affdaec3e; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.certificates
    ADD CONSTRAINT fk_rails_4affdaec3e FOREIGN KEY (course_id) REFERENCES public.courses(id);


--
-- Name: communities fk_rails_4b8f0bd893; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.communities
    ADD CONSTRAINT fk_rails_4b8f0bd893 FOREIGN KEY (school_id) REFERENCES public.schools(id);


--
-- Name: issued_certificates fk_rails_539d702198; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.issued_certificates
    ADD CONSTRAINT fk_rails_539d702198 FOREIGN KEY (certificate_id) REFERENCES public.certificates(id);


--
-- Name: school_links fk_rails_551faa9777; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.school_links
    ADD CONSTRAINT fk_rails_551faa9777 FOREIGN KEY (school_id) REFERENCES public.schools(id);


--
-- Name: user_activities fk_rails_56e545161c; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.user_activities
    ADD CONSTRAINT fk_rails_56e545161c FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: quizzes fk_rails_5f5df25460; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.quizzes
    ADD CONSTRAINT fk_rails_5f5df25460 FOREIGN KEY (target_id) REFERENCES public.targets(id);


--
-- Name: startup_feedback fk_rails_60705e3e6e; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.startup_feedback
    ADD CONSTRAINT fk_rails_60705e3e6e FOREIGN KEY (faculty_id) REFERENCES public.faculty(id);


--
-- Name: target_groups fk_rails_68e7c56696; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.target_groups
    ADD CONSTRAINT fk_rails_68e7c56696 FOREIGN KEY (level_id) REFERENCES public.levels(id);


--
-- Name: course_exports fk_rails_6b435e4656; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.course_exports
    ADD CONSTRAINT fk_rails_6b435e4656 FOREIGN KEY (course_id) REFERENCES public.courses(id);


--
-- Name: target_versions fk_rails_6cc72128ba; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.target_versions
    ADD CONSTRAINT fk_rails_6cc72128ba FOREIGN KEY (target_id) REFERENCES public.targets(id);


--
-- Name: posts fk_rails_70d0b6486a; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT fk_rails_70d0b6486a FOREIGN KEY (topic_id) REFERENCES public.topics(id);


--
-- Name: connect_requests fk_rails_747fd9ee5a; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.connect_requests
    ADD CONSTRAINT fk_rails_747fd9ee5a FOREIGN KEY (connect_slot_id) REFERENCES public.connect_slots(id);


--
-- Name: faculty_startup_enrollments fk_rails_7dd48e645b; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.faculty_startup_enrollments
    ADD CONSTRAINT fk_rails_7dd48e645b FOREIGN KEY (faculty_id) REFERENCES public.faculty(id);


--
-- Name: faculty_course_enrollments fk_rails_7f37a352c7; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.faculty_course_enrollments
    ADD CONSTRAINT fk_rails_7f37a352c7 FOREIGN KEY (faculty_id) REFERENCES public.faculty(id);


--
-- Name: target_resources fk_rails_8008d91332; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.target_resources
    ADD CONSTRAINT fk_rails_8008d91332 FOREIGN KEY (target_id) REFERENCES public.targets(id);


--
-- Name: quiz_questions fk_rails_8aeeccde35; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.quiz_questions
    ADD CONSTRAINT fk_rails_8aeeccde35 FOREIGN KEY (correct_answer_id) REFERENCES public.answer_options(id);


--
-- Name: startups fk_rails_9778445529; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.startups
    ADD CONSTRAINT fk_rails_9778445529 FOREIGN KEY (level_id) REFERENCES public.levels(id);


--
-- Name: faculty_startup_enrollments fk_rails_a2ecf6226c; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.faculty_startup_enrollments
    ADD CONSTRAINT fk_rails_a2ecf6226c FOREIGN KEY (startup_id) REFERENCES public.startups(id);


--
-- Name: school_admins fk_rails_ad83b2aa73; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.school_admins
    ADD CONSTRAINT fk_rails_ad83b2aa73 FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: courses fk_rails_adf7d91583; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.courses
    ADD CONSTRAINT fk_rails_adf7d91583 FOREIGN KEY (school_id) REFERENCES public.schools(id);


--
-- Name: community_course_connections fk_rails_b0d8aa657f; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.community_course_connections
    ADD CONSTRAINT fk_rails_b0d8aa657f FOREIGN KEY (community_id) REFERENCES public.communities(id);


--
-- Name: target_evaluation_criteria fk_rails_b117456f32; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.target_evaluation_criteria
    ADD CONSTRAINT fk_rails_b117456f32 FOREIGN KEY (target_id) REFERENCES public.targets(id);


--
-- Name: target_resources fk_rails_b515a4aa50; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.target_resources
    ADD CONSTRAINT fk_rails_b515a4aa50 FOREIGN KEY (resource_id) REFERENCES public.resources(id);


--
-- Name: timeline_event_files fk_rails_bedd6f7432; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.timeline_event_files
    ADD CONSTRAINT fk_rails_bedd6f7432 FOREIGN KEY (timeline_event_id) REFERENCES public.timeline_events(id);


--
-- Name: course_exports fk_rails_bf1979272f; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.course_exports
    ADD CONSTRAINT fk_rails_bf1979272f FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: active_storage_attachments fk_rails_c3b3935057; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.active_storage_attachments
    ADD CONSTRAINT fk_rails_c3b3935057 FOREIGN KEY (blob_id) REFERENCES public.active_storage_blobs(id);


--
-- Name: admin_users fk_rails_c4f75db4e4; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.admin_users
    ADD CONSTRAINT fk_rails_c4f75db4e4 FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: quiz_questions fk_rails_c723d3feef; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.quiz_questions
    ADD CONSTRAINT fk_rails_c723d3feef FOREIGN KEY (quiz_id) REFERENCES public.quizzes(id);


--
-- Name: school_strings fk_rails_cf2944d2e5; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.school_strings
    ADD CONSTRAINT fk_rails_cf2944d2e5 FOREIGN KEY (school_id) REFERENCES public.schools(id);


--
-- Name: domains fk_rails_cfd32a1ccc; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.domains
    ADD CONSTRAINT fk_rails_cfd32a1ccc FOREIGN KEY (school_id) REFERENCES public.schools(id);


--
-- Name: faculty_course_enrollments fk_rails_d118bb40e4; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.faculty_course_enrollments
    ADD CONSTRAINT fk_rails_d118bb40e4 FOREIGN KEY (course_id) REFERENCES public.courses(id);


--
-- Name: connect_slots fk_rails_d2126a0b97; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.connect_slots
    ADD CONSTRAINT fk_rails_d2126a0b97 FOREIGN KEY (faculty_id) REFERENCES public.faculty(id);


--
-- Name: founders fk_rails_d61facdaac; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.founders
    ADD CONSTRAINT fk_rails_d61facdaac FOREIGN KEY (college_id) REFERENCES public.colleges(id);


--
-- Name: course_authors fk_rails_dcac0ae302; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.course_authors
    ADD CONSTRAINT fk_rails_dcac0ae302 FOREIGN KEY (course_id) REFERENCES public.courses(id);


--
-- Name: levels fk_rails_e397b0035d; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.levels
    ADD CONSTRAINT fk_rails_e397b0035d FOREIGN KEY (course_id) REFERENCES public.courses(id);


--
-- Name: users fk_rails_e7d0538b2c; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT fk_rails_e7d0538b2c FOREIGN KEY (school_id) REFERENCES public.schools(id);


--
-- Name: timeline_events fk_rails_ec87f2b4bf; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.timeline_events
    ADD CONSTRAINT fk_rails_ec87f2b4bf FOREIGN KEY (evaluator_id) REFERENCES public.faculty(id);


--
-- Name: course_authors fk_rails_eec0f30346; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.course_authors
    ADD CONSTRAINT fk_rails_eec0f30346 FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: issued_certificates fk_rails_f999b73926; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.issued_certificates
    ADD CONSTRAINT fk_rails_f999b73926 FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: target_evaluation_criteria fk_rails_fa032f8917; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.target_evaluation_criteria
    ADD CONSTRAINT fk_rails_fa032f8917 FOREIGN KEY (evaluation_criterion_id) REFERENCES public.evaluation_criteria(id);


--
-- Name: applicants fk_rails_fcb549781c; Type: FK CONSTRAINT; Schema: public; Owner: mac
--

ALTER TABLE ONLY public.applicants
    ADD CONSTRAINT fk_rails_fcb549781c FOREIGN KEY (course_id) REFERENCES public.courses(id);


--
-- PostgreSQL database dump complete
--

