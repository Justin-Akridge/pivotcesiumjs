-- Create Companies Table
DO $$
BEGIN
  CREATE TABLE IF NOT EXISTS companies (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL UNIQUE,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
  );
END $$;

-- Create Users Table
DO $$
BEGIN
  CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    Name TEXT NOT NULL,
    company_id UUID NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    role TEXT NOT NULL CHECK (role IN ('user', 'admin')),
    permissions TEXT NOT NULL CHECK (permissions IN ('read', 'write')),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (company_id) REFERENCES companies(id)
  );
END $$;

-- Create Jobs Table
DO $$
BEGIN
  CREATE TABLE IF NOT EXISTS jobs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    company_id UUID NOT NULL,
    name TEXT NOT NULL,
    client_name TEXT NOT NULL,
    status TEXT DEFAULT 'pending' CHECK (status IN ('field', 'pending', 'completed')),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE CASCADE
  );
END $$;

-- Create Audit Table
-- Used for auditing who and what actions were done
DO $$
BEGIN
  CREATE TABLE IF NOT EXISTS job_audit (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    job_id UUID NOT NULL,
    modified_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    modified_by TEXT NOT NULL,
    action TEXT NOT NULL CHECK (action IN ('CREATE', 'UPDATE', 'DELETE')),
    changes TEXT,
    FOREIGN KEY (job_id) REFERENCES jobs(id) on DELETE CASCADE
  );
END $$;

DO $$
BEGIN
  CREATE TABLE IF NOT EXISTS file_metadata (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    filename TEXT NOT NULL,
    file_size BIGINT NOT NULL,
    upload_time TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    file_path TEXT NOT NULL,
    job_id UUID REFERENCES jobs(id) ON DELETE CASCADE,
    UNIQUE(filename, job_id)
  );
END $$;


-- Create Poles Table
DO $$
BEGIN
  CREATE TABLE IF NOT EXISTS poles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    job_id UUID NOT NULL,
    pole_data JSONB,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (job_id) REFERENCES jobs(id) ON DELETE CASCADE,
    file_metadata_id UUID REFERENCES file_metadata(id) ON DELETE CASCADE
  );
END $$;

-- Create Poles Table
DO $$
BEGIN
  CREATE TABLE IF NOT EXISTS poles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    job_id UUID NOT NULL,
    pole_data JSONB,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (job_id) REFERENCES jobs(id) ON DELETE CASCADE,
    file_metadata_id UUID REFERENCES file_metadata(id) ON DELETE CASCADE
  );
END $$;

-- Create Midspans Table
DO $$
BEGIN
  CREATE TABLE IF NOT EXISTS midspans (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    job_id UUID NOT NULL,
    midspan_data JSONB,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (job_id) REFERENCES jobs(id) ON DELETE CASCADE,
    file_metadata_id UUID REFERENCES file_metadata(id) ON DELETE CASCADE
  );
END $$;

-- Create Vegetation Table
DO $$
BEGIN
  CREATE TABLE IF NOT EXISTS vegetation (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    job_id UUID NOT NULL,
    vegetation_data JSONB,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (job_id) REFERENCES jobs(id) ON DELETE CASCADE,
    file_metadata_id UUID REFERENCES file_metadata(id) ON DELETE CASCADE
  );
END $$;


