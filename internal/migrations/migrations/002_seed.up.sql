-- Insert 48 services
INSERT INTO services (name, description)
VALUES ('Billing', 'Handles invoices and payments'),
       ('Auth', 'User authentication and tokens'),
       ('Notifications', 'Manages email, SMS, and push notifications'),
       ('Reporting', 'Generates analytics and business reports'),
       ('Search', 'Provides search functionality across services'),
       ('Security', 'Ensures data protection and access control'),
       ('User Profile', 'Manages user information and preferences'),
       ('File Storage', 'Handles file uploads and storage'),
       ('Image Processing', 'Optimizes and resizes images'),
       ('Video Streaming', 'Provides on-demand and live video services'),
       ('Chat', 'Enables real-time chat functionality'),
       ('Calendar', 'Handles events and scheduling'),
       ('Contacts', 'Manages user address books and contacts'),
       ('Payments', 'Processes transactions and refunds'),
       ('Subscriptions', 'Manages recurring billing and subscriptions'),
       ('Inventory', 'Keeps track of product stock levels'),
       ('Orders', 'Handles order creation and tracking'),
       ('Shipping', 'Manages deliveries and logistics'),
       ('CRM', 'Customer relationship management system'),
       ('Support Desk', 'Helpdesk ticketing and support'),
       ('Feedback', 'Collects customer ratings and reviews'),
       ('Marketing', 'Campaigns and promotions'),
       ('Email Service', 'Transactional and marketing emails'),
       ('SMS Gateway', 'Send and receive SMS messages'),
       ('Push Service', 'Mobile app push notifications'),
       ('IoT Hub', 'Manages IoT devices and telemetry'),
       ('AI Recommendations', 'Personalized recommendations engine'),
       ('Fraud Detection', 'Detects suspicious transactions'),
       ('Logging', 'Collects application logs'),
       ('Monitoring', 'Health checks and metrics'),
       ('Alerts', 'Triggers notifications based on rules'),
       ('Audit Trail', 'Keeps track of system activity'),
       ('Data Warehouse', 'Centralized analytics data store'),
       ('ETL Pipeline', 'Extract, transform, and load processes'),
       ('Machine Learning', 'Model training and inference'),
       ('Translation', 'Language translation services'),
       ('Localization', 'Support for multiple languages and regions'),
       ('Geo Service', 'Location-based data and maps'),
       ('Weather', 'Provides weather information'),
       ('News Feed', 'Aggregates and curates news'),
       ('Recommendations', 'Suggests products or content'),
       ('Personalization', 'Customizes user experiences'),
       ('Workflow', 'Manages tasks and automation'),
       ('Approvals', 'Approval flows and decisions'),
       ('DevOps Tools', 'CI/CD and developer utilities'),
       ('Test Service', 'Automated testing and QA'),
       ('Sandbox', 'Isolated test environments') ON CONFLICT (name) DO NOTHING;

-- Insert versions for all services
INSERT INTO versions (service_uuid, name, published_on)
SELECT s.service_uuid, v.name, v.published_on::date
FROM (VALUES
          -- Billing
          ('Billing', '2.1.0', '2025-07-10'),
          ('Billing', '2.0.5', '2025-06-01'),
          ('Billing', '1.9.9', '2025-05-15'),

          -- Auth
          ('Auth', '1.3.0', '2025-07-25'),
          ('Auth', '1.2.9', '2025-06-15'),
          ('Auth', '1.2.5', '2025-05-20'),

          -- Notifications
          ('Notifications', '3.0.0', '2025-07-20'),
          ('Notifications', '2.9.1', '2025-06-18'),
          ('Notifications', '2.8.5', '2025-05-10'),

          -- Reporting
          ('Reporting', '1.5.0', '2025-07-05'),
          ('Reporting', '1.4.8', '2025-06-10'),
          ('Reporting', '1.4.0', '2025-05-01'),

          -- Search
          ('Search', '0.9.2', '2025-07-12'),
          ('Search', '0.9.1', '2025-06-02'),
          ('Search', '0.8.9', '2025-05-05'),

          -- Security
          ('Security', '4.2.0', '2025-07-30'),
          ('Security', '4.1.5', '2025-06-22'),
          ('Security', '4.0.9', '2025-05-25'),

          -- User Profile
          ('User Profile', '1.1.0', '2025-07-15'),
          ('User Profile', '1.0.5', '2025-06-12'),
          ('User Profile', '1.0.0', '2025-05-08'),

          -- File Storage
          ('File Storage', '2.0.1', '2025-07-18'),
          ('File Storage', '2.0.0', '2025-06-11'),
          ('File Storage', '1.9.8', '2025-05-22'),

          -- Image Processing
          ('Image Processing', '3.1.0', '2025-07-28'),
          ('Image Processing', '3.0.5', '2025-06-15'),
          ('Image Processing', '2.9.9', '2025-05-30'),

          -- Video Streaming
          ('Video Streaming', '5.0.0', '2025-07-08'),
          ('Video Streaming', '4.9.2', '2025-06-20'),
          ('Video Streaming', '4.8.7', '2025-05-25'),

          -- Chat
          ('Chat', '2.2.1', '2025-07-06'),
          ('Chat', '2.2.0', '2025-06-05'),
          ('Chat', '2.1.9', '2025-05-10'),

          -- Calendar
          ('Calendar', '1.4.0', '2025-07-17'),
          ('Calendar', '1.3.8', '2025-06-08'),
          ('Calendar', '1.3.5', '2025-05-09'),

          -- Contacts
          ('Contacts', '1.0.2', '2025-07-13'),
          ('Contacts', '1.0.1', '2025-06-06'),
          ('Contacts', '1.0.0', '2025-05-12'),

          -- Payments
          ('Payments', '3.0.1', '2025-07-24'),
          ('Payments', '3.0.0', '2025-06-21'),
          ('Payments', '2.9.9', '2025-05-29'),

          -- Subscriptions
          ('Subscriptions', '2.1.0', '2025-07-14'),
          ('Subscriptions', '2.0.5', '2025-06-09'),
          ('Subscriptions', '2.0.0', '2025-05-14'),

          -- Inventory
          ('Inventory', '1.7.0', '2025-07-19'),
          ('Inventory', '1.6.8', '2025-06-13'),
          ('Inventory', '1.6.5', '2025-05-18'),

          -- Orders
          ('Orders', '2.5.0', '2025-07-27'),
          ('Orders', '2.4.9', '2025-06-16'),
          ('Orders', '2.4.5', '2025-05-23'),

          -- Shipping
          ('Shipping', '3.2.0', '2025-07-21'),
          ('Shipping', '3.1.9', '2025-06-17'),
          ('Shipping', '3.1.5', '2025-05-26'),

          -- CRM
          ('CRM', '4.0.0', '2025-07-26'),
          ('CRM', '3.9.8', '2025-06-14'),
          ('CRM', '3.9.5', '2025-05-27')) AS v(service_name, name, published_on)
         JOIN services s ON s.name = v.service_name;
