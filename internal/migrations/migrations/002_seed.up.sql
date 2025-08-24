INSERT INTO services (name, description)
VALUES ('Billing', 'Handles invoices and payments'),
       ('Auth', 'User authentication and tokens') ON CONFLICT (name) DO NOTHING;

INSERT INTO versions (service_uuid, name, published_on)
SELECT s.service_uuid, v.name, v.published_on::date
FROM (VALUES ('Billing', '2.1.0', '2025-07-10'),
             ('Billing', '2.0.5', '2025-06-01'),
             ('Auth', '1.3.0', '2025-07-25')) AS v(service_name, name, published_on)
         JOIN services s ON s.name = v.service_name;

