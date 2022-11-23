package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upEls, downEls)
}

var migrate = `
CREATE TABLE IF NOT EXISTS questions(
    id uuid,
    test_id int,
	theme_id int,
    value text,
    variants uuid[],
    answer uuid,
    time real,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS variants(
    id uuid,
    value text,
    PRIMARY KEY(id)
);

INSERT INTO questions(id, test_id, theme_id, value, variants, answer, time) VALUES('ab169a28-124f-454d-8c84-99e110f3b013', 1, 1, 'Что такое многочлен?', '{"dae7a8aa-6168-485d-96d6-188a5a9fbdaf", "a11366b7-6e7f-4d93-ba8d-4d68e423da31", "7744edaa-b34d-47cb-96cf-a66136406cd9"}', '7744edaa-b34d-47cb-96cf-a66136406cd9', 0.5);
INSERT INTO variants(id, value) VALUES('dae7a8aa-6168-485d-96d6-188a5a9fbdaf', 'Упорядоченный набор векторов в векторном пространстве');
INSERT INTO variants(id, value) VALUES('a11366b7-6e7f-4d93-ba8d-4d68e423da31', 'Алгебраическое выражение');
INSERT INTO variants(id, value) VALUES('7744edaa-b34d-47cb-96cf-a66136406cd9', 'Это сумма одночленов');

INSERT INTO questions(id, test_id, theme_id, value, variants, answer, time) VALUES('3e6273eb-4e86-44af-a920-07363afcb753', 2, 1, 'В какой степени многочлен не имеет решений в действительных числах?', '{"46edac97-7229-4fc9-bc0b-8ef00485aef0", "6e02caf6-39cf-4ab4-b2c7-738c8da1ed5e", "dfe7bce4-b644-4fc0-8e55-5f33f512a2c8"}', '46edac97-7229-4fc9-bc0b-8ef00485aef0', 1);
INSERT INTO variants(id, value) VALUES('46edac97-7229-4fc9-bc0b-8ef00485aef0', '>= 5');
INSERT INTO variants(id, value) VALUES('6e02caf6-39cf-4ab4-b2c7-738c8da1ed5e', '>5');
INSERT INTO variants(id, value) VALUES('dfe7bce4-b644-4fc0-8e55-5f33f512a2c8', '>=4');

INSERT INTO questions(id, test_id, theme_id, value, variants, answer, time) VALUES('dc42ff82-2888-4ac7-9caa-cd3c3bcee826', 3, 1, 'Решите уравнение x^2+1=0', '{"b00c3907-5970-4ca8-9a97-09db17bda3ca", "efd96f08-11ed-4551-a240-43d2274ee545", "baa582c6-62fe-4d88-b71e-7729f821f865"}', 'b00c3907-5970-4ca8-9a97-09db17bda3ca', 3);
INSERT INTO variants(id, value) VALUES('b00c3907-5970-4ca8-9a97-09db17bda3ca', 'x1=i, x2=-i');
INSERT INTO variants(id, value) VALUES('efd96f08-11ed-4551-a240-43d2274ee545', 'x1=2i, x2=-2i');
INSERT INTO variants(id, value) VALUES('baa582c6-62fe-4d88-b71e-7729f821f865', 'x1=0.5i, x2=-0.5i');




INSERT INTO questions(id, test_id, theme_id, value, variants, answer, time) VALUES('b10ac320-0d00-4e07-ad33-37b25c2e35b7', 1, 2, 'Что такое собственный вектор?', '{"f1ac703a-37c5-4872-8493-fcd35ce167f6", "920c6457-412b-49b4-ae70-3e11f2d6ab59", "7744edaa-b34d-47cb-96cf-a66136406cd9"}', '920c6457-412b-49b4-ae70-3e11f2d6ab59', 0.5);
INSERT INTO variants(id, value) VALUES('f1ac703a-37c5-4872-8493-fcd35ce167f6', 'Упорядоченный набор векторов в векторном пространстве');
INSERT INTO variants(id, value) VALUES('920c6457-412b-49b4-ae70-3e11f2d6ab59', 'Вектор, при применение к которому оператора получается тот же вектор');

INSERT INTO questions(id, test_id, theme_id, value, variants, answer, time) VALUES('d0a38f37-bc8f-47ee-87d0-07be0519d9fd', 2, 2, 'Что стоит на главной диагонали матрицы ЖНФ?', '{"17b71f45-91db-4a17-9757-c1d1d0d2db66", "96cd43ae-cf16-49ec-86b5-f0bdce6339d1", "7744edaa-b34d-47cb-96cf-a66136406cd9"}', '17b71f45-91db-4a17-9757-c1d1d0d2db66', 1);
INSERT INTO variants(id, value) VALUES('17b71f45-91db-4a17-9757-c1d1d0d2db66', 'Собственные значения оператора');
INSERT INTO variants(id, value) VALUES('96cd43ae-cf16-49ec-86b5-f0bdce6339d1', 'Собственные векторы оператора');

INSERT INTO questions(id, test_id, theme_id, value, variants, answer, time) VALUES('d4d1d909-8d42-4e4c-95e9-f4b81ee8b50e', 3, 2, 'Найти собственный вектор для оператора заданного матрицей {{0, 1, 1},{0, 1 , 0},{-1, 0, 2}}', '{"d7529ee2-aebd-4a63-92b0-bb24c726fde1", "30dafd1f-3788-4a8d-926d-e8ed049e3294", "457b5974-2120-4d63-a8da-d1e6596fb0f2"}', 'd7529ee2-aebd-4a63-92b0-bb24c726fde1', 3);
INSERT INTO variants(id, value) VALUES('d7529ee2-aebd-4a63-92b0-bb24c726fde1', '1');
INSERT INTO variants(id, value) VALUES('30dafd1f-3788-4a8d-926d-e8ed049e3294', '2');
INSERT INTO variants(id, value) VALUES('457b5974-2120-4d63-a8da-d1e6596fb0f2', '3');





INSERT INTO questions(id, test_id, theme_id, value, variants, answer, time) VALUES('745a3d11-4bc9-4d1a-974c-a9f8998fafb5', 1, 3, 'Как называется формула, связывающая коэффициенты многочлена и его корни?', '{"b587a980-d71d-45b7-b685-d2557969e8d7", "221f6f33-80f4-4371-adde-8a51bb518e64", "7e73ef72-5187-4c59-b39e-889cbab388a9"}', 'b587a980-d71d-45b7-b685-d2557969e8d7', 0.5);
INSERT INTO variants(id, value) VALUES('b587a980-d71d-45b7-b685-d2557969e8d7', 'Формула Виета');
INSERT INTO variants(id, value) VALUES('221f6f33-80f4-4371-adde-8a51bb518e64', 'Формула Лагранжа');
INSERT INTO variants(id, value) VALUES('7e73ef72-5187-4c59-b39e-889cbab388a9', 'Формула Байеса');


INSERT INTO questions(id, test_id, theme_id, value, variants, answer, time) VALUES('577d0fd0-5da8-4113-9e4c-e08912873d8f', 2, 3, 'Решите уравнение x^2-2x-3', '{"738dc8cd-f165-4b1c-88ae-99a67e6c0eaa", "de22230b-df8c-4a65-bb6a-2c0545ec91e2", "ebf5d4ed-20c2-4995-bd01-57e13dccd52e"}', '738dc8cd-f165-4b1c-88ae-99a67e6c0eaa', 1);
INSERT INTO variants(id, value) VALUES('738dc8cd-f165-4b1c-88ae-99a67e6c0eaa', 'x1=3, x2=-1');
INSERT INTO variants(id, value) VALUES('de22230b-df8c-4a65-bb6a-2c0545ec91e2', 'x1=2, x2=2');
INSERT INTO variants(id, value) VALUES('ebf5d4ed-20c2-4995-bd01-57e13dccd52e', 'x1=5, x2=1');

INSERT INTO questions(id, test_id, theme_id, value, variants, answer, time) VALUES('2c5ec7b8-5f63-4b3f-9ada-1bff2737d857', 3, 3, 'Решите уравнение x^4+1=0', '{"adcf39af-6497-4b22-afb3-854d5189d994", "ce9d4f8c-c614-442f-84a8-68041c094612", "baa582c6-62fe-4d88-b71e-7729f821f865"}', 'd950b9c3-9f60-4f98-9163-5dc8432d24d5', 1.5);
INSERT INTO variants(id, value) VALUES('adcf39af-6497-4b22-afb3-854d5189d994', 'x1=(√2 - √(2-4))/2, x2=(√2 - √(2-4))/2, x3=(-√2 + √(2-4))/2, x4=(-√2 - √(2-4))/2');
INSERT INTO variants(id, value) VALUES('ce9d4f8c-c614-442f-84a8-68041c094612', 'x1=(√3 - √(5-2))/2, x2=(√3 - √(5-2))/2, x3=(-√3 + √(5-2))/2, x4=(-√3 - √(5-2))/2');
INSERT INTO variants(id, value) VALUES('d950b9c3-9f60-4f98-9163-5dc8432d24d5', 'x1=(√4 - √(7-9))/2, x2=(√4 - √(7-9))/2, x3=(-√4 + √(7-9))/2, x4=(-√4 - √(7-9))/2');
`

func upEls(tx *sql.Tx) error {
	_, err := tx.Exec(migrate)
	if err != nil {
		return err
	}
	return nil
}

func downEls(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
