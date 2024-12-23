PGDMP      "                |            faztix    17.2    17.0 6    -           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                           false            .           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                           false            /           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                           false            0           1262    131072    faztix    DATABASE     |   CREATE DATABASE faztix WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LC_COLLATE = 'C' LC_CTYPE = 'en';
    DROP DATABASE faztix;
                     postgres    false                        2615    2200    public    SCHEMA        CREATE SCHEMA public;
    DROP SCHEMA public;
                     pg_database_owner    false            1           0    0    SCHEMA public    COMMENT     6   COMMENT ON SCHEMA public IS 'standard public schema';
                        pg_database_owner    false    4            �            1259    188417    cinema    TABLE     0  CREATE TABLE public.cinema (
    id integer NOT NULL,
    cinema_name character varying(255),
    price integer,
    location character varying(255),
    date date,
    "time" time without time zone,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone
);
    DROP TABLE public.cinema;
       public         heap r       postgres    false    4            �            1259    188416    cinema_id_seq    SEQUENCE     �   CREATE SEQUENCE public.cinema_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 $   DROP SEQUENCE public.cinema_id_seq;
       public               postgres    false    4    224            2           0    0    cinema_id_seq    SEQUENCE OWNED BY     ?   ALTER SEQUENCE public.cinema_id_seq OWNED BY public.cinema.id;
          public               postgres    false    223            �            1259    172044    movie    TABLE     �  CREATE TABLE public.movie (
    id integer NOT NULL,
    title character varying(100),
    image_movie character varying(255),
    genre character varying(100),
    release_date date,
    duration time without time zone,
    director character varying(255),
    cast_actor character varying(255),
    synopsis text,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone
);
    DROP TABLE public.movie;
       public         heap r       postgres    false    4            �            1259    196627    movie_cinema    TABLE     k   CREATE TABLE public.movie_cinema (
    id integer NOT NULL,
    movie_id integer,
    cinema_id integer
);
     DROP TABLE public.movie_cinema;
       public         heap r       postgres    false    4            �            1259    196626    movie_cinema_id_seq    SEQUENCE     �   CREATE SEQUENCE public.movie_cinema_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 *   DROP SEQUENCE public.movie_cinema_id_seq;
       public               postgres    false    4    226            3           0    0    movie_cinema_id_seq    SEQUENCE OWNED BY     K   ALTER SEQUENCE public.movie_cinema_id_seq OWNED BY public.movie_cinema.id;
          public               postgres    false    225            �            1259    172043    movie_id_seq    SEQUENCE     �   CREATE SEQUENCE public.movie_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.movie_id_seq;
       public               postgres    false    222    4            4           0    0    movie_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.movie_id_seq OWNED BY public.movie.id;
          public               postgres    false    221            �            1259    196644    orders    TABLE       CREATE TABLE public.orders (
    id integer NOT NULL,
    movie_cinema_id integer,
    quantity integer,
    total_price integer,
    date date,
    "time" time without time zone,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone
);
    DROP TABLE public.orders;
       public         heap r       postgres    false    4            �            1259    196643    orders_id_seq    SEQUENCE     �   CREATE SEQUENCE public.orders_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 $   DROP SEQUENCE public.orders_id_seq;
       public               postgres    false    228    4            5           0    0    orders_id_seq    SEQUENCE OWNED BY     ?   ALTER SEQUENCE public.orders_id_seq OWNED BY public.orders.id;
          public               postgres    false    227            �            1259    163867    user_credentials    TABLE     �   CREATE TABLE public.user_credentials (
    id integer NOT NULL,
    email character varying(255) NOT NULL,
    password character varying(255) NOT NULL
);
 $   DROP TABLE public.user_credentials;
       public         heap r       postgres    false    4            �            1259    163866    user_credentials_id_seq    SEQUENCE     �   CREATE SEQUENCE public.user_credentials_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 .   DROP SEQUENCE public.user_credentials_id_seq;
       public               postgres    false    4    218            6           0    0    user_credentials_id_seq    SEQUENCE OWNED BY     S   ALTER SEQUENCE public.user_credentials_id_seq OWNED BY public.user_credentials.id;
          public               postgres    false    217            �            1259    163878    users    TABLE     Z  CREATE TABLE public.users (
    id integer NOT NULL,
    firstname character varying(80),
    lastname character varying(80),
    phone_number character varying(18),
    image character varying(255),
    user_credentials_id integer NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone
);
    DROP TABLE public.users;
       public         heap r       postgres    false    4            �            1259    163877    users_id_seq    SEQUENCE     �   CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.users_id_seq;
       public               postgres    false    4    220            7           0    0    users_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;
          public               postgres    false    219            u           2604    188420 	   cinema id    DEFAULT     f   ALTER TABLE ONLY public.cinema ALTER COLUMN id SET DEFAULT nextval('public.cinema_id_seq'::regclass);
 8   ALTER TABLE public.cinema ALTER COLUMN id DROP DEFAULT;
       public               postgres    false    223    224    224            s           2604    172047    movie id    DEFAULT     d   ALTER TABLE ONLY public.movie ALTER COLUMN id SET DEFAULT nextval('public.movie_id_seq'::regclass);
 7   ALTER TABLE public.movie ALTER COLUMN id DROP DEFAULT;
       public               postgres    false    222    221    222            w           2604    196630    movie_cinema id    DEFAULT     r   ALTER TABLE ONLY public.movie_cinema ALTER COLUMN id SET DEFAULT nextval('public.movie_cinema_id_seq'::regclass);
 >   ALTER TABLE public.movie_cinema ALTER COLUMN id DROP DEFAULT;
       public               postgres    false    226    225    226            x           2604    196647 	   orders id    DEFAULT     f   ALTER TABLE ONLY public.orders ALTER COLUMN id SET DEFAULT nextval('public.orders_id_seq'::regclass);
 8   ALTER TABLE public.orders ALTER COLUMN id DROP DEFAULT;
       public               postgres    false    228    227    228            p           2604    163870    user_credentials id    DEFAULT     z   ALTER TABLE ONLY public.user_credentials ALTER COLUMN id SET DEFAULT nextval('public.user_credentials_id_seq'::regclass);
 B   ALTER TABLE public.user_credentials ALTER COLUMN id DROP DEFAULT;
       public               postgres    false    218    217    218            q           2604    163881    users id    DEFAULT     d   ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);
 7   ALTER TABLE public.users ALTER COLUMN id DROP DEFAULT;
       public               postgres    false    220    219    220            &          0    188417    cinema 
   TABLE DATA           h   COPY public.cinema (id, cinema_name, price, location, date, "time", created_at, updated_at) FROM stdin;
    public               postgres    false    224   �?       $          0    172044    movie 
   TABLE DATA           �   COPY public.movie (id, title, image_movie, genre, release_date, duration, director, cast_actor, synopsis, created_at, updated_at) FROM stdin;
    public               postgres    false    222   �@       (          0    196627    movie_cinema 
   TABLE DATA           ?   COPY public.movie_cinema (id, movie_id, cinema_id) FROM stdin;
    public               postgres    false    226   ,U       *          0    196644    orders 
   TABLE DATA           r   COPY public.orders (id, movie_cinema_id, quantity, total_price, date, "time", created_at, updated_at) FROM stdin;
    public               postgres    false    228   �V                  0    163867    user_credentials 
   TABLE DATA           ?   COPY public.user_credentials (id, email, password) FROM stdin;
    public               postgres    false    218   W       "          0    163878    users 
   TABLE DATA           z   COPY public.users (id, firstname, lastname, phone_number, image, user_credentials_id, created_at, updated_at) FROM stdin;
    public               postgres    false    220   QY       8           0    0    cinema_id_seq    SEQUENCE SET     ;   SELECT pg_catalog.setval('public.cinema_id_seq', 6, true);
          public               postgres    false    223            9           0    0    movie_cinema_id_seq    SEQUENCE SET     C   SELECT pg_catalog.setval('public.movie_cinema_id_seq', 101, true);
          public               postgres    false    225            :           0    0    movie_id_seq    SEQUENCE SET     <   SELECT pg_catalog.setval('public.movie_id_seq', 101, true);
          public               postgres    false    221            ;           0    0    orders_id_seq    SEQUENCE SET     ;   SELECT pg_catalog.setval('public.orders_id_seq', 1, true);
          public               postgres    false    227            <           0    0    user_credentials_id_seq    SEQUENCE SET     F   SELECT pg_catalog.setval('public.user_credentials_id_seq', 10, true);
          public               postgres    false    217            =           0    0    users_id_seq    SEQUENCE SET     ;   SELECT pg_catalog.setval('public.users_id_seq', 13, true);
          public               postgres    false    219            �           2606    188425    cinema cinema_pkey 
   CONSTRAINT     P   ALTER TABLE ONLY public.cinema
    ADD CONSTRAINT cinema_pkey PRIMARY KEY (id);
 <   ALTER TABLE ONLY public.cinema DROP CONSTRAINT cinema_pkey;
       public                 postgres    false    224            �           2606    196632    movie_cinema movie_cinema_pkey 
   CONSTRAINT     \   ALTER TABLE ONLY public.movie_cinema
    ADD CONSTRAINT movie_cinema_pkey PRIMARY KEY (id);
 H   ALTER TABLE ONLY public.movie_cinema DROP CONSTRAINT movie_cinema_pkey;
       public                 postgres    false    226            �           2606    172052    movie movie_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.movie
    ADD CONSTRAINT movie_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.movie DROP CONSTRAINT movie_pkey;
       public                 postgres    false    222            �           2606    196650    orders orders_pkey 
   CONSTRAINT     P   ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_pkey PRIMARY KEY (id);
 <   ALTER TABLE ONLY public.orders DROP CONSTRAINT orders_pkey;
       public                 postgres    false    228            {           2606    163876 +   user_credentials user_credentials_email_key 
   CONSTRAINT     g   ALTER TABLE ONLY public.user_credentials
    ADD CONSTRAINT user_credentials_email_key UNIQUE (email);
 U   ALTER TABLE ONLY public.user_credentials DROP CONSTRAINT user_credentials_email_key;
       public                 postgres    false    218            }           2606    163874 &   user_credentials user_credentials_pkey 
   CONSTRAINT     d   ALTER TABLE ONLY public.user_credentials
    ADD CONSTRAINT user_credentials_pkey PRIMARY KEY (id);
 P   ALTER TABLE ONLY public.user_credentials DROP CONSTRAINT user_credentials_pkey;
       public                 postgres    false    218                       2606    163884    users users_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pkey;
       public                 postgres    false    220            �           2606    163886 #   users users_user_credentials_id_key 
   CONSTRAINT     m   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_user_credentials_id_key UNIQUE (user_credentials_id);
 M   ALTER TABLE ONLY public.users DROP CONSTRAINT users_user_credentials_id_key;
       public                 postgres    false    220            �           2606    196638 (   movie_cinema movie_cinema_cinema_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.movie_cinema
    ADD CONSTRAINT movie_cinema_cinema_id_fkey FOREIGN KEY (cinema_id) REFERENCES public.cinema(id) ON DELETE CASCADE;
 R   ALTER TABLE ONLY public.movie_cinema DROP CONSTRAINT movie_cinema_cinema_id_fkey;
       public               postgres    false    4741    226    224            �           2606    196633 '   movie_cinema movie_cinema_movie_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.movie_cinema
    ADD CONSTRAINT movie_cinema_movie_id_fkey FOREIGN KEY (movie_id) REFERENCES public.movie(id) ON DELETE CASCADE;
 Q   ALTER TABLE ONLY public.movie_cinema DROP CONSTRAINT movie_cinema_movie_id_fkey;
       public               postgres    false    226    222    4739            �           2606    196651 "   orders orders_movie_cinema_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_movie_cinema_id_fkey FOREIGN KEY (movie_cinema_id) REFERENCES public.movie_cinema(id) ON DELETE CASCADE;
 L   ALTER TABLE ONLY public.orders DROP CONSTRAINT orders_movie_cinema_id_fkey;
       public               postgres    false    226    4743    228            �           2606    163887 $   users users_user_credentials_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_user_credentials_id_fkey FOREIGN KEY (user_credentials_id) REFERENCES public.user_credentials(id) ON DELETE CASCADE;
 N   ALTER TABLE ONLY public.users DROP CONSTRAINT users_user_credentials_id_fkey;
       public               postgres    false    4733    220    218            &   �   x���K�0�u{� H;�(���kl�6)���m�M4�f6�L��3�\�s��#V�;�?���#V�9�F�6��_�gE�
?�{20���M6�ʹ$՞B`_���Vy)���f4�	KWar��ИS�C�k���v��*O���p��FGl���Q1��Ќ|�F�5>��;\���AV	��	� s      $      x��]�r[7�}F��<͓����)_��i;����LW�l�dF�!)g���pH�v���E�j�c�D�`XX{aI���f��Ni��IΈڸ m�1�x�:�*�R�5-�?���1���-*�i�9[U���\H꡶�`�/O�-���,O�������pJ���wZ�û��+�r��������t(��r5�g��N������u͛�O������y�8�|�9�j�T\�A���?�Hቍb���;�����Z?�_�c�!K��ܚ�-X�M����{/�6�%��_�X?�B�!
�`?�=����Z.د��Q>Z�U�c��VmV��փ~���ӳ5�UO����z�����l=;>IB|�5�du�Q�$h��6��כ�(	_%��<7:i|��8���5B;U??����.1�{M��/4;�}q��_����A��f��_���l�����!��xA��	�y�qO�qS��1��Um!k�MQ����9�{Z��.71�F#ؓ���z�|������I����B�tGɲ�����3�8�|��]�<&ӄjx2Ip�|՘\Q�+�ο���PVtJ���Sf����UP��S]�*c��OOix�Y.�j�9���mh����}�zD4�5^��@�d9QӼ)_E��i��o���H�����FH��iy0�m�Zo��g��掲���e���^oV�f&Kg�,�yL-lH����$�b%[Jn_h��W��7�G��h쓒�M���%ǐ	�)�T�!�H��k�1r�c���D�;}d?�Ѫ�h��m ������,Ĥ#�������ǜ<��լ��\܍p8"Ƭğ��;J�ج��b�W��Ҽ���'L:��C��U���8�Q�Rc�Յ*�-PX
Ę���:(|W����rN;�YvZb�b��[�� "e�y,I;&��/����)_x�*+j�NOiu2�m1;~�a�h]��< ���	��\�l��:�E��Tr$�R"n@����$5{L̢�q=�-֌\�-7�+ō�9��U�_��*xѵ`�k$Ѳ��a���l؏g��,��e�7TY�����T+X�l�T�\S��,ˍ�gd�R#n��w����u@1`!;4/ձ1���r�YeG�HS���\�DXğ��ʑ1�ٜ�oY�?K+�S�Ŗ�O\)<�����.�}��"�,��U<�aU�ݽD�TWVZ�NŌ�/Pe*'�
lHI�h�C���I-�v$��!�2��U��4�Q���VP�{��<�����F)s��^X�#�q�ӉCސ���fyV��pT�M�6���4��l({.3	APU��
(92g_��;ʁ��lU���U=gM�Rt��
��;���V]��S@\��� oȉ�*5�=�T�-�s��Ic�M� u<HKt�O�-$�Lv�2�E�Fޫ� |W�W��cLQ~�r�x�F�x ����p6�x_�P���*��
x{W	�����x�>.�gg}��� C�hr͕W¶l�' �o�$x��M_qQ*��ʪ_q���M%�\$�ѻe����ue-Tatj���M���Zz���^���a����Wi�W:#���j2��=�^Jϖg+Vk��A��FGT�ZK. ۡ�H:��p$�� ���w���Ke�Q\��%�S�R���L���#+Vq]P���:5R��6�'��z�RJ&�\���V8Q)�Z�"�hAR}��j�Fޫ��8|g�z`��|�����h�q�[���5<���֖�M�u-�E��Fz�<�O��]mKq,_���9I�x2<��r<J��چRb*[������U@��zG9����'���9U�@=:�)��5�^�V2�7�:�<��]o[E��o�F�%{N���'L��- S+�$̽�CHh�Y&!4*��X�w�o-W`�u�A����ǳ՜�f���q���ӍQ��ESv��`�/�o����D��aj��C��r���0G�Y7� �u��7!O��A�Ѡtd�U�Ri���0��C��W*m>��7�~\5�^�����{c����
��~�[$��	-m�b�\�{Ů�ݟ�е�+�^�����|x(P���&��(�������mQ��H���6��A���d2�N��r��xzLǳ��c�H�j�&2a�X6K��5�TB���ޝڵĵE|{�	����\��kF��o|捐�T��ʭ�V&�(���r�(�qpz	�뷴8�O��Ɖ�5/�1�����6~�ǭYȭ&
��Џ3�8���w���������ծ�=�n%S��;�fה<��jx���h��'QyY���̐����]MQec򍸵���u@��b
j��$a?-����ti��Ё�*���RsԮ�w "@ǂZߊV�]�\�i�H�[],O��%
�/|؄��>jT���j��~��gw-o#�AR19��6ظ�S�U�H#`�1�_�d�G-��GE�X�r�Ժ4����}+k$�F"�I}����r�fgk&�5�	�lE��!�<�2���2��F־�5s��ߪ���aOiUKMd>��<�.�pn�録$���P�v�s�����ʇe����fE�ygu�
�kPkt�MJ�q�}0֑tyb��<��]���J*�2��o��^�[k�_Ip*�xFjW�OYبR�>�6,vqc�*���Z�q&��p�������ȕ%mc��L��aI9�C��G_s�Z��+6��o���!���	��3�`�璻���썬�b*E^��j���(�������ɹ?��a/����
�����|T�^�;��lʍ�Cz�jd�& Rha5�5D�Rx`kj����Ed,��B1��m�Y���&"C_157��p�jP�%f�XG��H��s�AT�>I̾�5�`+�Kba�~��b��Y�� �DL�H�^_d�B�of��J��,&g��3�Y�*��;i�PO������Ѣ*�P�~(���m�㰥�K���R���7��7��J��m����Com9
Շ�(����5ͭF|�2k?J�E�-f���:�Z�H�g��!ArT4�[/����ö�u�A���n��M��2�	�`:����xjiI�p�î�n-���C�����;�,Wx��5��x��찰�h��9��*/��26�]#�:���^ذ���fV���H�@\l�'|����1�~#�񵸯�5�m�p~�����D���8��E��W�s�Z��?R;%��!�J�k�eط���& ��b*gN��O0B
*��S��`J�)i�ҵ�z.�{Vۋ�{'�W��d%{���@CU}Ú�$L�P��X��\�7�[�������փ�,�m�c٤�U�jڧ4F�,O]db@iZn1���k����h�p����k�j}tq�w���$"�cM�'6�\ �[��[/��&��S᠔�1;�[�Ț�w%_u��� F��72$_D��u������;�%S�������)-֬VG�`W�`���Q:%�IT%�b)�]Sr}�j"������޻{Z[���~K�pǛ���Nc�Q�xT��^�+�[�{'��L�g?��z�i�O�嘛�
i"A
U�<������֮����&��^���'&�˳�ސ<���'�r�P�l���8�VU���z
��go��p���q�3�"���`)������c�XsF�R�B�~��Ķ���L8H9Ĕ��VU��C@��0��""E�HH��U�Q���mNz�����K��ق���5�0��u���^x���%�]�$��v����5M8|�TF{2_����S:�[��1�j���G9��Xd��YU�,g�+���ߢ��k���/�>z~������GO�X-E�w�غ�kFSi��Q-M�l�&��c��j01L�l�ߥ2�Pk�����3�b�5���C�d�N$WL����A�!&wkoٛ��^�����d���ډ��n%x(�n-g�����4ݸ�2��'��F��t����'R�>c�"��Z`��ڝ��J�`�'~�����O�~J"�G��~���
�Dq��9Y�j;%��<�B&�{OxLl��[߄�!~�LQ`?� {  ��b�]5L��9��XY��4���Iӯ#�H�z�9�k^o}�b*s���4�*�1<�T��9#�k�>q�-lMR�� ��]J�nh�w��( �fP����3mV�9S�J��xn:o-r�CMPFf�"*�[Y���+��½PB�~�m�i�L����W�5^v@�4	FF��g08^\��z,�
br����o�3�b:�`kV��R���v5�J蜄ʨ3��/�'��C��*!�f�v������kX/�'�G(��0��n�s]�	��L��B�`�sZ��t�u6[,X�םM´��O�jy0�q�R���ˊ�l+�5za��ͪ�� 6��s�{�|?��9���u��A�����H�BC�n���U^����!����ˋ�f9^�����@��:�*Eu��{y��̊���I��j������;?��VF�I��r���	'��Fe{�ތ�NE�]�;H��E��`��ud88��T�{<;��qǢH"��*B|�o��j#�(�[s�=qF�oL���a�$�2��|�fMd~�h�ۙ�tL��D)nH�nֻ��=��-���Dڢ�8m�bI�(��3�C�q<�� U(>�fۭ2.�ѷFߪb�B�z:Æ�]F��kR�mф=�dC�2�n�^/���#�����Q�7��o�R&5S�V��0/V��u�N{�e�ą�֑���Dv���6��U�.���[5�y8�PRy�j�^O����|k�pXb�I�0L����&�������K��JBbT�b�Sy�xbw�{�M8l��䤱췡��~��͆X�FgY�����9�W��RX_m����ܝ����uQ�7b�����d؛�Ey�\�GZ�� ����A7���ɲ����G�	ߜP"�9��f3�<��f������!f+-ҡ�2],A"��!p��*��)��Ķq��W88�D�"���i/��_W���z?�$���)^-��qX���Y6��&����.lз~
����_r��r^j߼Vy^Y5��f|��趆%p�4�f�}�B�L[qᅾ5M���
)��1m�_�e�(^r�z��-(q �IU� �Ȯ܊_x�o}Z]1�����w:���t�8f��D����ҝP��t �f�D�-$/{���S���9� t�?������լ      (   �  x��˕�0Cע�9�zy��1WY�ز��>ʘʩK���4�*Z�c��vx�����P]����4������"�zv���W���	�9�c��۴��SZΤτВ&����~��gW뫈0����4��ٚ�c���ݗW� �z�������t����n���1t_��i�5�����iI�}��3-Q,��Hp!��3��h�Z$R��6�߳��p+.����Wc��iO-wE�-Wr��f�f���3UĆ�\�s��K���!~|��{�X�Z��8�no	f���Y�'����
,�i�%;�s[ƣ�:��|�x�~�[�g�Ӳ�^�ˠ�8,���/Z�x�smo��cb]�w�	�7�ۼl�%���Û�D:~���p-      *   A   x�%��	�0��4 ��BTD*p�uؐ���4opz)�}�������w�Wv�
���D� ���          .  x���ˎ�@ @�5|;:�S�Fmh�!�),�YU���_?��|�ۻ9�S��|�{����H����O�G�r�=��ɡ��X(�]J�R�d��,�sL`U)�o(�W!'�l �~���t:L71�)�b�z#<V��H��c����d��S���!9j��h������P�����KL�o?QTv�|?�pm?���� �ۏ{19T�5HV����T�-n�]<NrF�|;�36%��\U�MG�0N{�Mlv��o��';EL����_���X��'��'�۝�
O v��s�����z�����j`�.	:�N�=��,���3�
��pmᨴ�+�1s���Z
�pO�
�oy���*�����L]A��wI���;7�OJ�л�NPPp4��ҊT��_�W�.ļm�k�XA�\�ٚ)o܄�`��"�xX�!��W�t�i�Z-.��D=kw/������w�٣B��2v�i���v���$270���H�tl2�i�VW~ֶ��!�D�!�F����ۗ�~2�n4�R&5�v|:����Ų�b�%�      "   .  x�U�9n�@Ek����f=DN��ڼ�#Br������,����0�>�`�f+ �f�����g `duĎ�e)$��>F8}4����D�P�8`�w.�\lig[G[��ΣAE��y| ��S�
���
0-�j0nWP	^c���qݦ����!�Eja�I��!j��9�&����//��Ƃ�����%�۵֟�`���I8�DO}pڛ9COn�����':�Q��Vu!s�|0�.�����1$�)��|��<�Ð��a��Y�1�}��"��$�<uM��k�u/     