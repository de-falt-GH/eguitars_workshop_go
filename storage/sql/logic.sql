------ ТРИГГЕР + ФУНКЦИЯ 1
CREATE OR REPLACE FUNCTION decrement_component() RETURNS TRIGGER AS
$$
BEGIN
    IF (TG_OP = 'INSERT') THEN
        IF (SELECT quantity FROM component WHERE id = NEW.component_id) < 1 THEN
            RAISE EXCEPTION 'Not enough % available', (SELECT "name" FROM component WHERE id  = NEW.id);
            RETURN NULL;
        ELSE
            UPDATE component SET quantity = quantity - 1 WHERE id = NEW.component_id;
        END IF;
    ELSIF (TG_OP = 'DELETE') THEN
        UPDATE component SET quantity = quantity + 1 WHERE id = OLD.component_id;
    END IF;
    RETURN NEW;
END; 
$$
LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER component_decrement_trigger
    BEFORE INSERT OR DELETE ON required_components
    FOR EACH ROW
    EXECUTE FUNCTION decrement_component();
------
    
------ ТРИГГЕР + ФУНКЦИЯ 2
CREATE OR REPLACE FUNCTION total_purchase_update() RETURNS TRIGGER AS
$$
BEGIN
    IF (TG_OP = 'INSERT') THEN
        UPDATE customer SET total_purchase = total_purchase + NEW.price WHERE id = NEW.id;
    ELSIF (TG_OP = 'UPDATE') THEN
        UPDATE customer SET total_purchase = total_purchase + NEW.price - OLD.price WHERE id = NEW.id;
    ELSIF (TG_OP = 'DELETE') THEN
        UPDATE customer SET total_purchase = total_purchase - OLD.price WHERE id = NEW.id;
    END IF;
    RETURN NULL;
END;
$$
LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER total_purchase_update_trigger
    AFTER INSERT OR DELETE OR UPDATE OF price ON "order"
    FOR EACH ROW
    EXECUTE FUNCTION total_purchase_update();
------

------ ТРИГГЕР + ФУНКЦИЯ 3
CREATE OR REPLACE FUNCTION customer_rank_update() RETURNS TRIGGER AS
$$
BEGIN
    IF (NEW.total_purchase < 10000) THEN
        UPDATE customer SET customer_rank_id = 1 WHERE id = NEW.id;
    ELSIF (NEW.total_purchase < 50000) THEN
        UPDATE customer SET customer_rank_id = 2 WHERE id = NEW.id;
    ELSIF (NEW.total_purchase < 150000) THEN
        UPDATE customer SET customer_rank_id = 3 WHERE id = NEW.id;
    ELSE
        UPDATE customer SET customer_rank_id = 4 WHERE id = NEW.id;
    END IF;
    RETURN NULL;
END;
$$
LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER customer_rank_update_trigger
    AFTER UPDATE OF total_purchase ON customer
    FOR EACH ROW
    EXECUTE FUNCTION customer_rank_update();
------

------ Процедура 1
CREATE OR REPLACE PROCEDURE drop_all_tables()
AS $$
BEGIN
    DROP TABLE IF EXISTS personal_info CASCADE;
    DROP TABLE IF EXISTS customer_rank CASCADE;
    DROP TABLE IF EXISTS customer CASCADE;
    DROP TABLE IF EXISTS master_rank CASCADE;
    DROP TABLE IF EXISTS "master" CASCADE;
    DROP TABLE IF EXISTS guitar CASCADE;
    DROP TABLE IF EXISTS component CASCADE;
    DROP TABLE IF EXISTS order_type CASCADE;
    DROP TABLE IF EXISTS order_status CASCADE;
    DROP TABLE IF EXISTS "order" CASCADE;
    DROP TABLE IF EXISTS required_components CASCADE;
    DROP TABLE IF EXISTS "user" CASCADE;
    DROP TABLE IF EXISTS permissions CASCADE;
    DROP TABLE IF EXISTS credentials CASCADE;
END;
$$
LANGUAGE plpgsql;
------

------ Процедура 2
CREATE OR REPLACE PROCEDURE change_order_master_by_id(order_id int, new_master_id int)
AS $$
BEGIN
    UPDATE "order" SET master_id = new_master_id WHERE id = order_id;
END;
$$
LANGUAGE plpgsql;
------

------ Процедура 3
CREATE OR REPLACE PROCEDURE promote_master(master_id int)
AS $$
BEGIN
    IF (SELECT master_rank_id FROM "master" WHERE id = master_id) < 4 THEN
        UPDATE master SET rank_id = rank_id + 1 WHERE id = master_id;
    ELSE
        RAISE EXCEPTION 'У УКАЗАННОГО МАСТЕРА УЖЕ МАКСИМАЛЬНЫЙ РАНГ';
    END IF;
END;
$$
LANGUAGE plpgsql;
------

