PGDMP                     	    |            bpjstk    15.8 (Homebrew)    15.8 (Homebrew)     p           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            q           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            r           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            s           1262    32776    bpjstk    DATABASE     h   CREATE DATABASE bpjstk WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'C';
    DROP DATABASE bpjstk;
                hasbi    false                        2615    32847    soal1    SCHEMA        CREATE SCHEMA soal1;
    DROP SCHEMA soal1;
                hasbi    false                        2615    32777    soal2    SCHEMA        CREATE SCHEMA soal2;
    DROP SCHEMA soal2;
                hasbi    false                        3079    32778 	   uuid-ossp 	   EXTENSION     ?   CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;
    DROP EXTENSION "uuid-ossp";
                   false            t           0    0    EXTENSION "uuid-ossp"    COMMENT     W   COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';
                        false    2            �            1255    32802    get_children_data(jsonb)    FUNCTION     �  CREATE FUNCTION public.get_children_data(datajson jsonb) RETURNS jsonb
    LANGUAGE plpgsql
    AS $$
DECLARE
    result JSONB;
BEGIN
    WITH RECURSIVE child_tree AS (
        SELECT id, name_data, parent_id
        FROM soal2.user_data
        WHERE id = (dataJson->>'uuidParent')::UUID

        UNION ALL

        SELECT u.id, u.name_data, u.parent_id
        FROM soal2.user_data u
        INNER JOIN child_tree ct ON u.parent_id = ct.id  -- Join with the CTE to find children
    )
    SELECT jsonb_agg(child_tree) INTO result
    FROM child_tree
    WHERE id <> (dataJson->>'uuidParent')::UUID;  
    RETURN jsonb_build_object('data', result);
END;
$$;
 8   DROP FUNCTION public.get_children_data(datajson jsonb);
       public          hasbi    false            �            1255    41125    create_child_data(jsonb)    FUNCTION     �  CREATE FUNCTION soal1.create_child_data(input_json jsonb) RETURNS character varying
    LANGUAGE plpgsql
    AS $$
DECLARE

BEGIN
    -- Loop through each parent data entry in the JSON

        -- Insert into the parent table and get the inserted ID
        INSERT INTO soal1.data ( start_date, end_date, nominal)
        VALUES (
            (input_json->>'startDate')::DATE,  -- Cast to DATE
            (input_json->>'endDate')::DATE,    -- Cast to DATE
            (input_json->>'nominal')::numeric,
            (input_json->>'parent_id')::uuid
        );  -- Capture the inserted parent ID
    RETURN 'Success';
EXCEPTION
    WHEN OTHERS THEN
        RETURN 'Error: ' || SQLERRM;  -- Return the error message
END;
$$;
 9   DROP FUNCTION soal1.create_child_data(input_json jsonb);
       soal1          hasbir    false    8            �            1255    41121    create_parent_data(jsonb)    FUNCTION     �  CREATE FUNCTION soal1.create_parent_data(input_json jsonb) RETURNS character varying
    LANGUAGE plpgsql
    AS $$
DECLARE
    parent_id UUID;  -- Variable to hold the inserted parent ID
 BEGIN

      INSERT INTO soal1.data ( start_date, end_date, nominal)
      VALUES (
           (input_json->>'startDate')::DATE,  -- Cast to DATE
           (input_json->>'endDate')::DATE,    -- Cast to DATE
           (input_json->>'nominal')::NUMERIC
            )

        RETURNING id INTO parent_id;  -- Capture the inserted parent ID

      

    RETURN parent_id;
EXCEPTION
    WHEN OTHERS THEN
        RETURN 'Error: ' || SQLERRM;  -- Return the error message
END;
$$;
 :   DROP FUNCTION soal1.create_parent_data(input_json jsonb);
       soal1          hasbir    false    8            �            1255    32927 #   get_all_data_with_pagination(jsonb)    FUNCTION     �  CREATE FUNCTION soal1.get_all_data_with_pagination(datas jsonb) RETURNS jsonb
    LANGUAGE plpgsql
    AS $$
DECLARE
    result JSONB;
    p_page INT := COALESCE((datas->>'page')::INT, 1); 
    p_limit INT := COALESCE((datas->>'limit')::INT, 100); 
    p_offset INT := (p_page - 1) * p_limit;  
BEGIN
    SELECT jsonb_agg(row_to_json(d.*))
    INTO result
    FROM (
        SELECT 
            id,
            nominal,
            start_date AS startDate,
            end_date AS endDate
        FROM soal1."data"
        ORDER BY created_at DESC -- Adjust the order as needed
        OFFSET p_offset LIMIT p_limit
    ) AS d;

    RETURN COALESCE(result, '[]'::jsonb); -- Return an empty JSONB array if no rows found
END;
$$;
 ?   DROP FUNCTION soal1.get_all_data_with_pagination(datas jsonb);
       soal1          hasbi    false    8            �            1255    32931    get_child_data(jsonb)    FUNCTION     �  CREATE FUNCTION soal1.get_child_data(datajson jsonb) RETURNS jsonb
    LANGUAGE plpgsql
    AS $$
DECLARE
    result JSONB;
BEGIN
    -- Selecting and aggregating the data into JSONB, sorting within the aggregate
    SELECT jsonb_agg(
        jsonb_build_object(
            'id', id,
            'startDate', start_date,
            'nominal', nominal,
            'endDate', end_date,
            'parentId', parent_id
        ) ORDER BY start_date
    ) INTO result
    FROM soal1.data_detail s
    WHERE s.parent_id = (datajson->>'parentId')::uuid;

    -- Returning the result as a JSONB object
    RETURN COALESCE(result, '[]'::jsonb); -- Return empty array if no result
END;
$$;
 4   DROP FUNCTION soal1.get_child_data(datajson jsonb);
       soal1          hasbi    false    8            �            1255    32868    insert_bulk_data(jsonb)    FUNCTION       CREATE FUNCTION soal1.insert_bulk_data(input_json jsonb) RETURNS character varying
    LANGUAGE plpgsql
    AS $$
DECLARE
    data_record JSONB;  -- Variable for each parent data record
    parent_id UUID;  -- Variable to hold the inserted parent ID
    child_record JSONB;  -- Variable for each child record
BEGIN
    -- Loop through each parent data entry in the JSON
    FOR data_record IN SELECT * FROM jsonb_array_elements(input_json->'data') LOOP
        -- RAISE NOTICE for parent data before insertion

        -- Insert into the parent table and get the inserted ID
        INSERT INTO soal1.data ( start_date, end_date, nominal)
        VALUES (
            (data_record->'parentData'->>'startDate')::DATE,  -- Cast to DATE
            (data_record->'parentData'->>'endDate')::DATE,    -- Cast to DATE
            (data_record->'parentData'->>'nominal')::NUMERIC
        )
        RETURNING id INTO parent_id;  -- Capture the inserted parent ID

        -- Raise notice for parent insertion
        RAISE NOTICE 'Inserted parent record with ID: %', parent_id;

        -- Loop through each child data entry and insert into the child table
        FOR child_record IN SELECT * FROM jsonb_array_elements(data_record->'childData') LOOP


            -- Insert into the child table
            INSERT INTO soal1.data_detail (start_date, end_date, nominal, parent_id)
            VALUES (
                (child_record->>'startDate')::DATE,  -- Cast to DATE
                (child_record->>'endDate')::DATE,    -- Cast to DATE
                (child_record->>'nominal')::NUMERIC,
                parent_id
                
                  -- Set the parent_id to the captured ID
            );

            -- Raise notice for child insertion
            RAISE NOTICE 'Inserted child record with nominal: % for parent ID: %', 
                (child_record->>'nominal'), parent_id;
        END LOOP;
    END LOOP;

    RETURN 'Success';
EXCEPTION
    WHEN OTHERS THEN
        RETURN 'Error: ' || SQLERRM;  -- Return the error message
END;
$$;
 8   DROP FUNCTION soal1.insert_bulk_data(input_json jsonb);
       soal1          hasbi    false    8            �            1255    32804    get_all_data(jsonb)    FUNCTION     �  CREATE FUNCTION soal2.get_all_data(datajson jsonb) RETURNS jsonb
    LANGUAGE plpgsql
    AS $$
DECLARE
    result JSONB;
BEGIN

    SELECT jsonb_agg(
        jsonb_build_object(
            'id', id,
            'nameData', name_data,
            'parentId', parent_id
        )
    ) INTO result
    FROM soal2.user_data;

    -- Returning the result as a JSONB object
    RETURN result;
END;
$$;
 2   DROP FUNCTION soal2.get_all_data(datajson jsonb);
       soal2          hasbi    false    7            �            1255    32803    get_children_data(jsonb)    FUNCTION     ,  CREATE FUNCTION soal2.get_children_data(datajson jsonb) RETURNS jsonb
    LANGUAGE plpgsql
    AS $$
DECLARE
    result JSONB;
BEGIN
    WITH RECURSIVE child_tree AS (
        SELECT id, name_data AS nameData, parent_id AS parentId
        FROM soal2.user_data
        WHERE id = (dataJson->>'uuidParent')::UUID

        UNION ALL

        SELECT u.id, u.name_data AS nameData, u.parent_id AS parentId
        FROM soal2.user_data u
        INNER JOIN child_tree ct ON u.parent_id = ct.id  -- Join with the CTE to find children
    )

   
    SELECT jsonb_agg(
        jsonb_build_object(
            'id', id,
            'nameData', ct.nameData,
           'parentId', ct.parentId
        )
    ) INTO result
    FROM child_tree ct
    where id <> (dataJson->>'uuidParent')::UUID;

    RETURN result;
END;
$$;
 7   DROP FUNCTION soal2.get_children_data(datajson jsonb);
       soal2          hasbi    false    7            �            1259    49322    data    TABLE     �   CREATE TABLE soal1.data (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    start_date date NOT NULL,
    end_date date NOT NULL,
    nominal numeric(15,2) NOT NULL
);
    DROP TABLE soal1.data;
       soal1         heap    hasbir    false    2    8            �            1259    49329    data_detail    TABLE       CREATE TABLE soal1.data_detail (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    start_date date NOT NULL,
    end_date date NOT NULL,
    nominal numeric(15,2) NOT NULL,
    parent_id uuid NOT NULL
);
    DROP TABLE soal1.data_detail;
       soal1         heap    hasbir    false    2    8            �            1259    32789 	   user_data    TABLE     �   CREATE TABLE soal2.user_data (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    name_data character varying(255) NOT NULL,
    parent_id uuid
);
    DROP TABLE soal2.user_data;
       soal2         heap    hasbi    false    2    7            l          0    49322    data 
   TABLE DATA           L   COPY soal1.data (id, created_at, start_date, end_date, nominal) FROM stdin;
    soal1          hasbir    false    218   `4       m          0    49329    data_detail 
   TABLE DATA           ^   COPY soal1.data_detail (id, created_at, start_date, end_date, nominal, parent_id) FROM stdin;
    soal1          hasbir    false    219   }4       k          0    32789 	   user_data 
   TABLE DATA           <   COPY soal2.user_data (id, name_data, parent_id) FROM stdin;
    soal2          hasbi    false    217   �4       �           2606    49335    data_detail data_detail_pkey 
   CONSTRAINT     Y   ALTER TABLE ONLY soal1.data_detail
    ADD CONSTRAINT data_detail_pkey PRIMARY KEY (id);
 E   ALTER TABLE ONLY soal1.data_detail DROP CONSTRAINT data_detail_pkey;
       soal1            hasbir    false    219            �           2606    49328    data data_pkey 
   CONSTRAINT     K   ALTER TABLE ONLY soal1.data
    ADD CONSTRAINT data_pkey PRIMARY KEY (id);
 7   ALTER TABLE ONLY soal1.data DROP CONSTRAINT data_pkey;
       soal1            hasbir    false    218            �           2606    32796 !   user_data user_data_name_data_key 
   CONSTRAINT     `   ALTER TABLE ONLY soal2.user_data
    ADD CONSTRAINT user_data_name_data_key UNIQUE (name_data);
 J   ALTER TABLE ONLY soal2.user_data DROP CONSTRAINT user_data_name_data_key;
       soal2            hasbi    false    217            �           2606    32794    user_data user_data_pkey 
   CONSTRAINT     U   ALTER TABLE ONLY soal2.user_data
    ADD CONSTRAINT user_data_pkey PRIMARY KEY (id);
 A   ALTER TABLE ONLY soal2.user_data DROP CONSTRAINT user_data_pkey;
       soal2            hasbi    false    217            �           2606    32797 "   user_data user_data_parent_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY soal2.user_data
    ADD CONSTRAINT user_data_parent_id_fkey FOREIGN KEY (parent_id) REFERENCES soal2.user_data(id) ON DELETE SET NULL;
 K   ALTER TABLE ONLY soal2.user_data DROP CONSTRAINT user_data_parent_id_fkey;
       soal2          hasbi    false    217    3543    217            l      x������ � �      m      x������ � �      k      x������ � �     