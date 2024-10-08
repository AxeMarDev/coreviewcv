toc.dat                                                                                             0000600 0004000 0002000 00000035737 14640031431 0014452 0                                                                                                    ustar 00postgres                        postgres                        0000000 0000000                                                                                                                                                                        PGDMP                           |           coreview    14.11 (Homebrew)    14.11 (Homebrew) 3    z           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false         {           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false         |           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false         }           1262    24601    coreview    DATABASE     S   CREATE DATABASE coreview WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'C';
    DROP DATABASE coreview;
                axellmartinez    false         �            1259    24642    client    TABLE     �   CREATE TABLE public.client (
    id integer NOT NULL,
    name character varying(255),
    username character varying(50),
    email character varying(255),
    hash_password text,
    phone character varying(15),
    company_id integer
);
    DROP TABLE public.client;
       public         heap    axellmartinez    false         �            1259    24641    client_id_seq    SEQUENCE     �   CREATE SEQUENCE public.client_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 $   DROP SEQUENCE public.client_id_seq;
       public          axellmartinez    false    216         ~           0    0    client_id_seq    SEQUENCE OWNED BY     ?   ALTER SEQUENCE public.client_id_seq OWNED BY public.client.id;
          public          axellmartinez    false    215         �            1259    24603    company    TABLE     �   CREATE TABLE public.company (
    id integer NOT NULL,
    company_name character varying(255),
    company_code character varying(50),
    masteremployee_id integer
);
    DROP TABLE public.company;
       public         heap    axellmartinez    false         �            1259    24602    company_id_seq    SEQUENCE     �   CREATE SEQUENCE public.company_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 %   DROP SEQUENCE public.company_id_seq;
       public          axellmartinez    false    210                    0    0    company_id_seq    SEQUENCE OWNED BY     A   ALTER SEQUENCE public.company_id_seq OWNED BY public.company.id;
          public          axellmartinez    false    209         �            1259    24612    employee    TABLE       CREATE TABLE public.employee (
    id integer NOT NULL,
    name character varying(255),
    username character varying(50),
    email character varying(255),
    hash_password text,
    phone character varying(15),
    isadmin boolean,
    company_id integer
);
    DROP TABLE public.employee;
       public         heap    axellmartinez    false         �            1259    24611    employee_id_seq    SEQUENCE     �   CREATE SEQUENCE public.employee_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 &   DROP SEQUENCE public.employee_id_seq;
       public          axellmartinez    false    212         �           0    0    employee_id_seq    SEQUENCE OWNED BY     C   ALTER SEQUENCE public.employee_id_seq OWNED BY public.employee.id;
          public          axellmartinez    false    211         �            1259    24630    project    TABLE     r   CREATE TABLE public.project (
    id integer NOT NULL,
    name character varying(255),
    company_id integer
);
    DROP TABLE public.project;
       public         heap    axellmartinez    false         �            1259    24659    project_client    TABLE     h   CREATE TABLE public.project_client (
    project_id integer NOT NULL,
    client_id integer NOT NULL
);
 "   DROP TABLE public.project_client;
       public         heap    axellmartinez    false         �            1259    24674    project_employee    TABLE     l   CREATE TABLE public.project_employee (
    project_id integer NOT NULL,
    employee_id integer NOT NULL
);
 $   DROP TABLE public.project_employee;
       public         heap    axellmartinez    false         �            1259    24629    project_id_seq    SEQUENCE     �   CREATE SEQUENCE public.project_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 %   DROP SEQUENCE public.project_id_seq;
       public          axellmartinez    false    214         �           0    0    project_id_seq    SEQUENCE OWNED BY     A   ALTER SEQUENCE public.project_id_seq OWNED BY public.project.id;
          public          axellmartinez    false    213         �           2604    24645 	   client id    DEFAULT     f   ALTER TABLE ONLY public.client ALTER COLUMN id SET DEFAULT nextval('public.client_id_seq'::regclass);
 8   ALTER TABLE public.client ALTER COLUMN id DROP DEFAULT;
       public          axellmartinez    false    215    216    216         �           2604    24606 
   company id    DEFAULT     h   ALTER TABLE ONLY public.company ALTER COLUMN id SET DEFAULT nextval('public.company_id_seq'::regclass);
 9   ALTER TABLE public.company ALTER COLUMN id DROP DEFAULT;
       public          axellmartinez    false    209    210    210         �           2604    24615    employee id    DEFAULT     j   ALTER TABLE ONLY public.employee ALTER COLUMN id SET DEFAULT nextval('public.employee_id_seq'::regclass);
 :   ALTER TABLE public.employee ALTER COLUMN id DROP DEFAULT;
       public          axellmartinez    false    212    211    212         �           2604    24633 
   project id    DEFAULT     h   ALTER TABLE ONLY public.project ALTER COLUMN id SET DEFAULT nextval('public.project_id_seq'::regclass);
 9   ALTER TABLE public.project ALTER COLUMN id DROP DEFAULT;
       public          axellmartinez    false    214    213    214         u          0    24642    client 
   TABLE DATA           ]   COPY public.client (id, name, username, email, hash_password, phone, company_id) FROM stdin;
    public          axellmartinez    false    216       3701.dat o          0    24603    company 
   TABLE DATA           T   COPY public.company (id, company_name, company_code, masteremployee_id) FROM stdin;
    public          axellmartinez    false    210       3695.dat q          0    24612    employee 
   TABLE DATA           h   COPY public.employee (id, name, username, email, hash_password, phone, isadmin, company_id) FROM stdin;
    public          axellmartinez    false    212       3697.dat s          0    24630    project 
   TABLE DATA           7   COPY public.project (id, name, company_id) FROM stdin;
    public          axellmartinez    false    214       3699.dat v          0    24659    project_client 
   TABLE DATA           ?   COPY public.project_client (project_id, client_id) FROM stdin;
    public          axellmartinez    false    217       3702.dat w          0    24674    project_employee 
   TABLE DATA           C   COPY public.project_employee (project_id, employee_id) FROM stdin;
    public          axellmartinez    false    218       3703.dat �           0    0    client_id_seq    SEQUENCE SET     <   SELECT pg_catalog.setval('public.client_id_seq', 49, true);
          public          axellmartinez    false    215         �           0    0    company_id_seq    SEQUENCE SET     =   SELECT pg_catalog.setval('public.company_id_seq', 28, true);
          public          axellmartinez    false    209         �           0    0    employee_id_seq    SEQUENCE SET     >   SELECT pg_catalog.setval('public.employee_id_seq', 28, true);
          public          axellmartinez    false    211         �           0    0    project_id_seq    SEQUENCE SET     <   SELECT pg_catalog.setval('public.project_id_seq', 8, true);
          public          axellmartinez    false    213         �           2606    24653    client client_email_key 
   CONSTRAINT     S   ALTER TABLE ONLY public.client
    ADD CONSTRAINT client_email_key UNIQUE (email);
 A   ALTER TABLE ONLY public.client DROP CONSTRAINT client_email_key;
       public            axellmartinez    false    216         �           2606    24649    client client_pkey 
   CONSTRAINT     P   ALTER TABLE ONLY public.client
    ADD CONSTRAINT client_pkey PRIMARY KEY (id);
 <   ALTER TABLE ONLY public.client DROP CONSTRAINT client_pkey;
       public            axellmartinez    false    216         �           2606    24651    client client_username_key 
   CONSTRAINT     Y   ALTER TABLE ONLY public.client
    ADD CONSTRAINT client_username_key UNIQUE (username);
 D   ALTER TABLE ONLY public.client DROP CONSTRAINT client_username_key;
       public            axellmartinez    false    216         �           2606    24610     company company_company_code_key 
   CONSTRAINT     c   ALTER TABLE ONLY public.company
    ADD CONSTRAINT company_company_code_key UNIQUE (company_code);
 J   ALTER TABLE ONLY public.company DROP CONSTRAINT company_company_code_key;
       public            axellmartinez    false    210         �           2606    24608    company company_pkey 
   CONSTRAINT     R   ALTER TABLE ONLY public.company
    ADD CONSTRAINT company_pkey PRIMARY KEY (id);
 >   ALTER TABLE ONLY public.company DROP CONSTRAINT company_pkey;
       public            axellmartinez    false    210         �           2606    24623    employee employee_email_key 
   CONSTRAINT     W   ALTER TABLE ONLY public.employee
    ADD CONSTRAINT employee_email_key UNIQUE (email);
 E   ALTER TABLE ONLY public.employee DROP CONSTRAINT employee_email_key;
       public            axellmartinez    false    212         �           2606    24619    employee employee_pkey 
   CONSTRAINT     T   ALTER TABLE ONLY public.employee
    ADD CONSTRAINT employee_pkey PRIMARY KEY (id);
 @   ALTER TABLE ONLY public.employee DROP CONSTRAINT employee_pkey;
       public            axellmartinez    false    212         �           2606    24621    employee employee_username_key 
   CONSTRAINT     ]   ALTER TABLE ONLY public.employee
    ADD CONSTRAINT employee_username_key UNIQUE (username);
 H   ALTER TABLE ONLY public.employee DROP CONSTRAINT employee_username_key;
       public            axellmartinez    false    212         �           2606    24663 "   project_client project_client_pkey 
   CONSTRAINT     s   ALTER TABLE ONLY public.project_client
    ADD CONSTRAINT project_client_pkey PRIMARY KEY (project_id, client_id);
 L   ALTER TABLE ONLY public.project_client DROP CONSTRAINT project_client_pkey;
       public            axellmartinez    false    217    217         �           2606    24678 &   project_employee project_employee_pkey 
   CONSTRAINT     y   ALTER TABLE ONLY public.project_employee
    ADD CONSTRAINT project_employee_pkey PRIMARY KEY (project_id, employee_id);
 P   ALTER TABLE ONLY public.project_employee DROP CONSTRAINT project_employee_pkey;
       public            axellmartinez    false    218    218         �           2606    24635    project project_pkey 
   CONSTRAINT     R   ALTER TABLE ONLY public.project
    ADD CONSTRAINT project_pkey PRIMARY KEY (id);
 >   ALTER TABLE ONLY public.project DROP CONSTRAINT project_pkey;
       public            axellmartinez    false    214         �           2606    24654    client client_company_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.client
    ADD CONSTRAINT client_company_id_fkey FOREIGN KEY (company_id) REFERENCES public.company(id);
 G   ALTER TABLE ONLY public.client DROP CONSTRAINT client_company_id_fkey;
       public          axellmartinez    false    210    3528    216         �           2606    24689 &   company company_masteremployee_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.company
    ADD CONSTRAINT company_masteremployee_id_fkey FOREIGN KEY (masteremployee_id) REFERENCES public.employee(id);
 P   ALTER TABLE ONLY public.company DROP CONSTRAINT company_masteremployee_id_fkey;
       public          axellmartinez    false    212    210    3532         �           2606    24624 !   employee employee_company_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.employee
    ADD CONSTRAINT employee_company_id_fkey FOREIGN KEY (company_id) REFERENCES public.company(id);
 K   ALTER TABLE ONLY public.employee DROP CONSTRAINT employee_company_id_fkey;
       public          axellmartinez    false    212    3528    210         �           2606    24669 ,   project_client project_client_client_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.project_client
    ADD CONSTRAINT project_client_client_id_fkey FOREIGN KEY (client_id) REFERENCES public.client(id);
 V   ALTER TABLE ONLY public.project_client DROP CONSTRAINT project_client_client_id_fkey;
       public          axellmartinez    false    217    3540    216         �           2606    24664 -   project_client project_client_project_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.project_client
    ADD CONSTRAINT project_client_project_id_fkey FOREIGN KEY (project_id) REFERENCES public.project(id);
 W   ALTER TABLE ONLY public.project_client DROP CONSTRAINT project_client_project_id_fkey;
       public          axellmartinez    false    3536    217    214         �           2606    24636    project project_company_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.project
    ADD CONSTRAINT project_company_id_fkey FOREIGN KEY (company_id) REFERENCES public.company(id);
 I   ALTER TABLE ONLY public.project DROP CONSTRAINT project_company_id_fkey;
       public          axellmartinez    false    214    3528    210         �           2606    24684 2   project_employee project_employee_employee_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.project_employee
    ADD CONSTRAINT project_employee_employee_id_fkey FOREIGN KEY (employee_id) REFERENCES public.employee(id);
 \   ALTER TABLE ONLY public.project_employee DROP CONSTRAINT project_employee_employee_id_fkey;
       public          axellmartinez    false    3532    212    218         �           2606    24679 1   project_employee project_employee_project_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.project_employee
    ADD CONSTRAINT project_employee_project_id_fkey FOREIGN KEY (project_id) REFERENCES public.project(id);
 [   ALTER TABLE ONLY public.project_employee DROP CONSTRAINT project_employee_project_id_fkey;
       public          axellmartinez    false    214    3536    218                                         3701.dat                                                                                            0000600 0004000 0002000 00000003512 14640031431 0014241 0                                                                                                    ustar 00postgres                        postgres                        0000000 0000000                                                                                                                                                                        1	1	1	1	$2a$10$VEmzAVmW1QRdE/dKIUKwYux03ZnCmPdoNGsH3R1pxRoSLJbrA8GY2	1	1
11	sdfdsf	dvvdwf	fsd	$2a$10$s2J8c0hZWB7alrRp8uU1P.ISEYxXDngQHV2XoXr4RcT6ydi9n7DKy	wer	14
12				$2a$10$tEbuY0CH60MgMd/1RIFlh.e7B3XnqaNjsS.MxgLCCcj093d9OEIDe		14
13	32	23	23	$2a$10$Ootjh5LAUB3Dbgjhymkwierlj04FBzTrXMJtGMZl9/JjI6dfVPrxW	232	14
14	wef	wefwef	qwe	$2a$10$hp/d0EqJaCh6vqlhswZ9J.bwmbLswO5S38IHxvFo/TsPPsw5St2/m	wef	14
15	u	u	u	$2a$10$wN8lOH/dRAFcie8ar97ct.yqBZXaWiaz57TrJckpkpE/EmDQuew5S	u	14
16	we	we	we	$2a$10$UcU6I.YQel3C/vh.Ef7Ng.ZwuzvUq44FVg5Go8HmJ.L/ys4Tx.a2C	we	14
17	32	234	23423	$2a$10$RGrRHcrgRqNP09dD1YhWJu.JSB3jYKHrTJaN8EoJISKECatABWs.m	4324	14
18	32er	234erer	23423er	$2a$10$hR08xdcpkuasMldcqIQEguCC3QAyP/tLPWzj7WplyxHs3FTMRBIwi	4324er	\N
19	er	erer	er	$2a$10$lJtLxSg8Lzou1UfoiocSl.RT46tt0zNQzqUP5f.ynRre7oIVk76Bi	er	\N
22	23	223	232	$2a$10$pZVDAffX8s3OrqNcR99Fauz153zwsluR0l7.HO8th.7asitDr1/5W	323	\N
24	weddwew	weddwwwe	weweweew	$2a$10$cEHrzVy9d8NMDGrQquUpJeE6NOxiyZ/ya5YmB54ApkZ.4ncHh979y	wefwef	18
31	wefwef4	wefweft4	wefwef4r	$2a$10$jmmJ7hog1gBKKVPJMdYF9u.B.gl0qFsdJqPsPiVUf50OC4mHeUtyq	g4tg45t4	20
36	new	new	new	$2a$10$xIE8Pgm3sXan4TNEbKQQ0./9PyIGHVzqPeVqQLxf56AYZOuZAiDIu	new	20
42	Axell Martinez	Axemar	axellmartinez@gmail.com	$2a$10$UhWsHblH27KqqKgAh5x0neTdmjrdG1szDwXaqWh1B.KGXj5cFWG3C	(956)372-9814	19
44	ijoijerfr	wef	werfew	$2a$10$55UZoTxWQYgjQ4GjRNOFKeZkatkB55Sa/YM/q1rvt5fspQc9exKB.	rfwerrf	22
45	axell	axemar12	1212	$2a$10$qpQWsfOZaqJYUXiZ.tmNyuYJZRG3LfBN4FyMb.JJzJxyzJnQGjnCW	12	23
46	new5	new5	new5	$2a$10$W/A0PUSa7U/u.imT5u9V5eKEG6LFEo98tnoOwrcOq5VjkjTff0Fw2	new5	24
47	John Wayne	johnwyn	axell@axell.com	$2a$10$5Dg/ojuOYRg3pfcsRkhTWuXfhQiRkKXB6ZnjvqFjV0UaItDBuvXVi	n/a	25
48	the4	the4	the4	$2a$10$TGQQlfYSy9OVFxD7gVYtFuMRtraeI2KXjdsAUEmXPgIsoLG3Sv21W	the4	24
49	the5	the5	the5	$2a$10$0j.4pvgDccWS/q46AbYGneb3YhL.CyrMICzNvELrqud5IA1XPUb8u	the5	24
\.


                                                                                                                                                                                      3695.dat                                                                                            0000600 0004000 0002000 00000000626 14640031431 0014260 0                                                                                                    ustar 00postgres                        postgres                        0000000 0000000                                                                                                                                                                        1	ed	ed	\N
2	qewdfrgfher	frddwefrgtew	\N
3	ewrf	ewrfew	\N
5	ewrfrv	ewrfewrv	\N
6	Axell Electric	AXELLE	\N
7	wef	wefw	\N
8	wefwwfwef	fwefwee	\N
9	company	company	\N
10	axellelectric	axell	\N
11	me	me	\N
12	we	we	\N
13	de	de	\N
14	23	242	\N
15	r	r	\N
17	edk	edk	\N
18	po	po	\N
19	tg	tg	16
20	4r	4r	17
21	company1	company1	20
22	company2	company2	21
23	user3	user3	23
24	the	the	24
25	mee	mee	26
26			28
\.


                                                                                                          3697.dat                                                                                            0000600 0004000 0002000 00000004700 14640031431 0014257 0                                                                                                    ustar 00postgres                        postgres                        0000000 0000000                                                                                                                                                                        1	ADMIN	eded	eded	\N	eed	t	1
2	ADMIN	rewdefrdgtdrwdefrd	qdfrgtdwaefsrdgtf	\N	ewqrt	t	2
3	ADMIN	ewrfewrfrv	rfewrfrv	$2a$10$yjV4hosrwBX2Zrs5Ze10OOaH.I3YIf6WgmDGZP74l/tgTNW6o8oVG	werfwerv	t	5
4	ADMIN	axellmar	axellmartinez@gmail.com	$2a$10$1HqkPrzMgmOs9Iuosg4yuOID2ixneNiQP3zu1PQNrBiiNBOlmFoXq	999	t	6
5	ADMIN	wefwef	efwef	$2a$10$YzZ/2sHQb6UozXVWsVfBB.UYTaBuoOEC6zzmDWPtQ62iOkQt6v89O	wefwef	t	7
6	ADMIN	fwf	wefwefwef	$2a$10$X5cPQ5W0DsnG8gmCEJy43OSmMfFlQw6TOpm1uZh7L1LQfwLfUBSHy	wfwef	t	8
7	ADMIN	employee	1	$2a$10$hhy9.ysspdNYPfCzCvZ17u2Crg6ca7xmXYRNBmuzf.tZSvARFTm/K	12	t	9
8	ADMIN	axemar	email	$2a$10$9RAZaNS0QasKql9AbibjqeV1YAor0arpUG5CMdl.KloxE4xGerwea	na	t	10
9	ADMIN	me	me	$2a$10$shyEhxq.BBlwZggtJH/zCe9skmfxsotpXp1/XAQ/Xy/9DoFQB/JK.	me	t	11
10	ADMIN	we	we	$2a$10$JQ.dg7C5t7ttxgjzDEDUzetr4pswSpl1z4oGDcqxBc.nxe1Q3S8Dy	we	t	12
11	ADMIN	de	de	$2a$10$..kVE/3XBl9XPrO/Ur8FK.zYIY063q6xGZhOstCiCTKhl.kwZHBB.	de	t	13
12	ADMIN	23423	34234	$2a$10$Ht5hCDAj2hib7N8vh8DgAeGkSbavkXv/MeTJlojVnkB4ibkmx/Hp6	42342	t	14
13	ADMIN	r	r	$2a$10$QCVCshCNh66RiH3xmEIHrejlTvIW/uqHnRE3CMP75ao1JsXzspRDq	r	t	15
14	ADMIN	edk	edk	$2a$10$5Ru2FRYtiypgJ0FQ68zRWOLIn7e4VXsoj6Yy4fWpSCF7cky2b44Ly	edk	t	17
15	ADMIN	po	po	$2a$10$UjrvqfTKLWNwxTZKIGcoJ.E/3Wd3PMkpEFHl0MsWS52kxhU59iypa	po	t	18
16	ADMIN	tg	tg	$2a$10$mdydXxLiJPVvnYnpz4zr0eKtGVvasmsUOAuMoh6FLy6Cr5CKBXyAW	tg	t	19
17	ADMIN	4r	4r	$2a$10$gs7FzC4HMJ44aUpmRrkcpeO92V5VuPZjnsBIr5xP4WGsmU9xl7EWi	4r4	t	20
18	John Denver	JOHNDEN	johndenver@gmail.com	$2a$10$5RlyfGEhmr8uCdsFfSIwweNDZMKIsXf01zQOWV1rNez0kr4O8bre.	n/a	\N	19
19	username1	username1	username1	$2a$10$GvsgjLCbjT/tUVRR8Hqd9.nP.eqxJogvkIy8.F/h1aTn9Ito2byEy	username1	\N	19
20	ADMIN	company1	company1	$2a$10$CPwTRGipY8/k7Yva9vQ1yuQhFPdGQEZOkIayv2q06FoI9OLVMYi4.	company1	t	21
21	ADMIN	company2	company2	$2a$10$EgmwcY4WEynkoDC.Y0ZoQeQsgzOMsEf1lK5oInWU1gO/Qaf8.aDpm	company2	t	22
22	employee3	employee3	employee3	$2a$10$0BxYBs5AaEEToI0osGtMeu1s8dJyIBIp9SAnprN05Ey2ommeRcwxG	employee3	\N	22
23	ADMIN	user3	user3	$2a$10$zXVDnzXIXfT6mULWIdG2c.hi060XQaBJJ.ru.4dd/MFMu7anrCL0G	user3	t	23
24	ADMIN	the	the	$2a$10$40nh0ojfOj/ICRL9wMDrxOYw4UPFo6AAZRxdilX04LPWcRmU82NpC	the	t	24
25	the2	the2	the2	$2a$10$gMiKn8OutZmaUNkuk.6wce20gQSxupaXwq53DBELe1W.21MOiAQoG	the2	\N	24
26	ADMIN	mee	mee	$2a$10$waqjAlIdWEUOWrOfm7uY2efIkVV5MSfr24VC56BTXUhwDKfTN8M2y	mee	t	25
27	mee2	mee2	mee2	$2a$10$Upi7xGq09KrZE2apvd339O3UYxl.4SjVO0ppLdhzFKlxLMHqrXKwO	mee2	\N	25
28	ADMIN			$2a$10$yqO3pgN0S2m4XF6eTOaHJ.ze7daIRyq.TlxPPtU1jslj.NETRms/u		t	26
\.


                                                                3699.dat                                                                                            0000600 0004000 0002000 00000000172 14640031431 0014260 0                                                                                                    ustar 00postgres                        postgres                        0000000 0000000                                                                                                                                                                        1	23	19
2	23we	19
3	new project	19
4	axell	20
5	new project	19
6	new project	24
7	new project	25
8	6213 tecate dr	25
\.


                                                                                                                                                                                                                                                                                                                                                                                                      3702.dat                                                                                            0000600 0004000 0002000 00000000005 14640031431 0014234 0                                                                                                    ustar 00postgres                        postgres                        0000000 0000000                                                                                                                                                                        \.


                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           3703.dat                                                                                            0000600 0004000 0002000 00000000005 14640031431 0014235 0                                                                                                    ustar 00postgres                        postgres                        0000000 0000000                                                                                                                                                                        \.


                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           restore.sql                                                                                         0000600 0004000 0002000 00000027412 14640031431 0015366 0                                                                                                    ustar 00postgres                        postgres                        0000000 0000000                                                                                                                                                                        --
-- NOTE:
--
-- File paths need to be edited. Search for $$PATH$$ and
-- replace it with the path to the directory containing
-- the extracted data files.
--
--
-- PostgreSQL database dump
--

-- Dumped from database version 14.11 (Homebrew)
-- Dumped by pg_dump version 14.11 (Homebrew)

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

DROP DATABASE coreview;
--
-- Name: coreview; Type: DATABASE; Schema: -; Owner: axellmartinez
--

CREATE DATABASE coreview WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'C';


ALTER DATABASE coreview OWNER TO axellmartinez;

\connect coreview

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: client; Type: TABLE; Schema: public; Owner: axellmartinez
--

CREATE TABLE public.client (
    id integer NOT NULL,
    name character varying(255),
    username character varying(50),
    email character varying(255),
    hash_password text,
    phone character varying(15),
    company_id integer
);


ALTER TABLE public.client OWNER TO axellmartinez;

--
-- Name: client_id_seq; Type: SEQUENCE; Schema: public; Owner: axellmartinez
--

CREATE SEQUENCE public.client_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.client_id_seq OWNER TO axellmartinez;

--
-- Name: client_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: axellmartinez
--

ALTER SEQUENCE public.client_id_seq OWNED BY public.client.id;


--
-- Name: company; Type: TABLE; Schema: public; Owner: axellmartinez
--

CREATE TABLE public.company (
    id integer NOT NULL,
    company_name character varying(255),
    company_code character varying(50),
    masteremployee_id integer
);


ALTER TABLE public.company OWNER TO axellmartinez;

--
-- Name: company_id_seq; Type: SEQUENCE; Schema: public; Owner: axellmartinez
--

CREATE SEQUENCE public.company_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.company_id_seq OWNER TO axellmartinez;

--
-- Name: company_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: axellmartinez
--

ALTER SEQUENCE public.company_id_seq OWNED BY public.company.id;


--
-- Name: employee; Type: TABLE; Schema: public; Owner: axellmartinez
--

CREATE TABLE public.employee (
    id integer NOT NULL,
    name character varying(255),
    username character varying(50),
    email character varying(255),
    hash_password text,
    phone character varying(15),
    isadmin boolean,
    company_id integer
);


ALTER TABLE public.employee OWNER TO axellmartinez;

--
-- Name: employee_id_seq; Type: SEQUENCE; Schema: public; Owner: axellmartinez
--

CREATE SEQUENCE public.employee_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.employee_id_seq OWNER TO axellmartinez;

--
-- Name: employee_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: axellmartinez
--

ALTER SEQUENCE public.employee_id_seq OWNED BY public.employee.id;


--
-- Name: project; Type: TABLE; Schema: public; Owner: axellmartinez
--

CREATE TABLE public.project (
    id integer NOT NULL,
    name character varying(255),
    company_id integer
);


ALTER TABLE public.project OWNER TO axellmartinez;

--
-- Name: project_client; Type: TABLE; Schema: public; Owner: axellmartinez
--

CREATE TABLE public.project_client (
    project_id integer NOT NULL,
    client_id integer NOT NULL
);


ALTER TABLE public.project_client OWNER TO axellmartinez;

--
-- Name: project_employee; Type: TABLE; Schema: public; Owner: axellmartinez
--

CREATE TABLE public.project_employee (
    project_id integer NOT NULL,
    employee_id integer NOT NULL
);


ALTER TABLE public.project_employee OWNER TO axellmartinez;

--
-- Name: project_id_seq; Type: SEQUENCE; Schema: public; Owner: axellmartinez
--

CREATE SEQUENCE public.project_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.project_id_seq OWNER TO axellmartinez;

--
-- Name: project_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: axellmartinez
--

ALTER SEQUENCE public.project_id_seq OWNED BY public.project.id;


--
-- Name: client id; Type: DEFAULT; Schema: public; Owner: axellmartinez
--

ALTER TABLE ONLY public.client ALTER COLUMN id SET DEFAULT nextval('public.client_id_seq'::regclass);


--
-- Name: company id; Type: DEFAULT; Schema: public; Owner: axellmartinez
--

ALTER TABLE ONLY public.company ALTER COLUMN id SET DEFAULT nextval('public.company_id_seq'::regclass);


--
-- Name: employee id; Type: DEFAULT; Schema: public; Owner: axellmartinez
--

ALTER TABLE ONLY public.employee ALTER COLUMN id SET DEFAULT nextval('public.employee_id_seq'::regclass);


--
-- Name: project id; Type: DEFAULT; Schema: public; Owner: axellmartinez
--

ALTER TABLE ONLY public.project ALTER COLUMN id SET DEFAULT nextval('public.project_id_seq'::regclass);


--
-- Data for Name: client; Type: TABLE DATA; Schema: public; Owner: axellmartinez
--

COPY public.client (id, name, username, email, hash_password, phone, company_id) FROM stdin;
\.
COPY public.client (id, name, username, email, hash_password, phone, company_id) FROM '$$PATH$$/3701.dat';

--
-- Data for Name: company; Type: TABLE DATA; Schema: public; Owner: axellmartinez
--

COPY public.company (id, company_name, company_code, masteremployee_id) FROM stdin;
\.
COPY public.company (id, company_name, company_code, masteremployee_id) FROM '$$PATH$$/3695.dat';

--
-- Data for Name: employee; Type: TABLE DATA; Schema: public; Owner: axellmartinez
--

COPY public.employee (id, name, username, email, hash_password, phone, isadmin, company_id) FROM stdin;
\.
COPY public.employee (id, name, username, email, hash_password, phone, isadmin, company_id) FROM '$$PATH$$/3697.dat';

--
-- Data for Name: project; Type: TABLE DATA; Schema: public; Owner: axellmartinez
--

COPY public.project (id, name, company_id) FROM stdin;
\.
COPY public.project (id, name, company_id) FROM '$$PATH$$/3699.dat';

--
-- Data for Name: project_client; Type: TABLE DATA; Schema: public; Owner: axellmartinez
--

COPY public.project_client (project_id, client_id) FROM stdin;
\.
COPY public.project_client (project_id, client_id) FROM '$$PATH$$/3702.dat';

--
-- Data for Name: project_employee; Type: TABLE DATA; Schema: public; Owner: axellmartinez
--

COPY public.project_employee (project_id, employee_id) FROM stdin;
\.
COPY public.project_employee (project_id, employee_id) FROM '$$PATH$$/3703.dat';

--
-- Name: client_id_seq; Type: SEQUENCE SET; Schema: public; Owner: axellmartinez
--

SELECT pg_catalog.setval('public.client_id_seq', 49, true);


--
-- Name: company_id_seq; Type: SEQUENCE SET; Schema: public; Owner: axellmartinez
--

SELECT pg_catalog.setval('public.company_id_seq', 28, true);


--
-- Name: employee_id_seq; Type: SEQUENCE SET; Schema: public; Owner: axellmartinez
--

SELECT pg_catalog.setval('public.employee_id_seq', 28, true);


--
-- Name: project_id_seq; Type: SEQUENCE SET; Schema: public; Owner: axellmartinez
--

SELECT pg_catalog.setval('public.project_id_seq', 8, true);


--
-- Name: client client_email_key; Type: CONSTRAINT; Schema: public; Owner: axellmartinez
--

ALTER TABLE ONLY public.client
    ADD CONSTRAINT client_email_key UNIQUE (email);


--
-- Name: client client_pkey; Type: CONSTRAINT; Schema: public; Owner: axellmartinez
--

ALTER TABLE ONLY public.client
    ADD CONSTRAINT client_pkey PRIMARY KEY (id);


--
-- Name: client client_username_key; Type: CONSTRAINT; Schema: public; Owner: axellmartinez
--

ALTER TABLE ONLY public.client
    ADD CONSTRAINT client_username_key UNIQUE (username);


--
-- Name: company company_company_code_key; Type: CONSTRAINT; Schema: public; Owner: axellmartinez
--

ALTER TABLE ONLY public.company
    ADD CONSTRAINT company_company_code_key UNIQUE (company_code);


--
-- Name: company company_pkey; Type: CONSTRAINT; Schema: public; Owner: axellmartinez
--

ALTER TABLE ONLY public.company
    ADD CONSTRAINT company_pkey PRIMARY KEY (id);


--
-- Name: employee employee_email_key; Type: CONSTRAINT; Schema: public; Owner: axellmartinez
--

ALTER TABLE ONLY public.employee
    ADD CONSTRAINT employee_email_key UNIQUE (email);


--
-- Name: employee employee_pkey; Type: CONSTRAINT; Schema: public; Owner: axellmartinez
--

ALTER TABLE ONLY public.employee
    ADD CONSTRAINT employee_pkey PRIMARY KEY (id);


--
-- Name: employee employee_username_key; Type: CONSTRAINT; Schema: public; Owner: axellmartinez
--

ALTER TABLE ONLY public.employee
    ADD CONSTRAINT employee_username_key UNIQUE (username);


--
-- Name: project_client project_client_pkey; Type: CONSTRAINT; Schema: public; Owner: axellmartinez
--

ALTER TABLE ONLY public.project_client
    ADD CONSTRAINT project_client_pkey PRIMARY KEY (project_id, client_id);


--
-- Name: project_employee project_employee_pkey; Type: CONSTRAINT; Schema: public; Owner: axellmartinez
--

ALTER TABLE ONLY public.project_employee
    ADD CONSTRAINT project_employee_pkey PRIMARY KEY (project_id, employee_id);


--
-- Name: project project_pkey; Type: CONSTRAINT; Schema: public; Owner: axellmartinez
--

ALTER TABLE ONLY public.project
    ADD CONSTRAINT project_pkey PRIMARY KEY (id);


--
-- Name: client client_company_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: axellmartinez
--

ALTER TABLE ONLY public.client
    ADD CONSTRAINT client_company_id_fkey FOREIGN KEY (company_id) REFERENCES public.company(id);


--
-- Name: company company_masteremployee_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: axellmartinez
--

ALTER TABLE ONLY public.company
    ADD CONSTRAINT company_masteremployee_id_fkey FOREIGN KEY (masteremployee_id) REFERENCES public.employee(id);


--
-- Name: employee employee_company_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: axellmartinez
--

ALTER TABLE ONLY public.employee
    ADD CONSTRAINT employee_company_id_fkey FOREIGN KEY (company_id) REFERENCES public.company(id);


--
-- Name: project_client project_client_client_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: axellmartinez
--

ALTER TABLE ONLY public.project_client
    ADD CONSTRAINT project_client_client_id_fkey FOREIGN KEY (client_id) REFERENCES public.client(id);


--
-- Name: project_client project_client_project_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: axellmartinez
--

ALTER TABLE ONLY public.project_client
    ADD CONSTRAINT project_client_project_id_fkey FOREIGN KEY (project_id) REFERENCES public.project(id);


--
-- Name: project project_company_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: axellmartinez
--

ALTER TABLE ONLY public.project
    ADD CONSTRAINT project_company_id_fkey FOREIGN KEY (company_id) REFERENCES public.company(id);


--
-- Name: project_employee project_employee_employee_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: axellmartinez
--

ALTER TABLE ONLY public.project_employee
    ADD CONSTRAINT project_employee_employee_id_fkey FOREIGN KEY (employee_id) REFERENCES public.employee(id);


--
-- Name: project_employee project_employee_project_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: axellmartinez
--

ALTER TABLE ONLY public.project_employee
    ADD CONSTRAINT project_employee_project_id_fkey FOREIGN KEY (project_id) REFERENCES public.project(id);


--
-- PostgreSQL database dump complete
--

                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      